/// <reference path="./IPage.ts" />
import Ractive from "ractive";
import "../scss/move.scss";
import Button from "../views/Button";
import { handleResult } from "./commandHandler";
import PartyStatus from "../views/PartyStatus";
import { numberFormat } from "../models/numberFormat";

declare function sendCommand(c: string, data?: any): ICommandResult;

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
      },
    });
  }

  private send(command: string, data: any) {
    const result = sendCommand(command, data);
    handleResult(result, this.app, this.ractive);
  }
}
