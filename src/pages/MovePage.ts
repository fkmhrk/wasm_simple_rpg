/// <reference path="./IPage.ts" />
import Ractive from "ractive";
import "../scss/move.scss";
import Button from "../views/Button";
import { handleResult } from "./commandHandler";
import PartyStatus from "../views/PartyStatus";

declare function sendCommand(c: string, data: string): ICommandResult;

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
      on: {
        start: () => this.start(),
      },
    });
  }

  private start() {
    let data: string;
    try {
      data = localStorage.getItem("d") ?? "";
    } catch {
      data = "";
    }
    const result = sendCommand("start", data);
    handleResult(result, this.app, this.ractive);
  }
}
