/// <reference path="./IPage.ts" />
import Ractive from "ractive";
import "../scss/move.scss";
import Button from "../views/Button";
import { handleResult } from "./commandHandler";
import PartyStatus from "../views/PartyStatus";
import { numberFormat } from "../models/numberFormat";

declare function sendCommand(c: string, data?: any): ICommandResult;
declare function save(): { data: string; iv: string };

export default class MovePage implements IPage {
  private app: IApplication;
  ractive!: Ractive;

  constructor(app: IApplication) {
    this.app = app;
  }

  async onCreate() {
    const t = await this.app.fetchTemplate("move.html");
    this.ractive = new Ractive({
      el: "#container",
      template: t,
      components: {
        Button: Button,
        PartyStatus: PartyStatus,
      },
      data: {
        format: numberFormat,
      },
      on: {
        send: (e: any, command: string, data: any) => this.send(command, data),
        save: () => this.save(),
      },
    });
  }

  private send(command: string, data: any) {
    const result = sendCommand(command, data);
    handleResult(result, this.app, this.ractive);
  }

  private save() {
    const result = save();
    if (result["data"] != null && result["iv"] != null) {
      try {
        localStorage.setItem("d", result["data"]);
        localStorage.setItem("i", result["iv"]);
        alert("saved!");
      } catch {
        console.log("Failed to save");
      }
    }
  }
}
