/// <reference path="./IPage.ts" />
import Ractive from "ractive";
import "../scss/battle.scss";
import Button from "../views/Button";
import { handleResult } from "./commandHandler";
import PartyStatus from "../views/PartyStatus";
import { numberFormat } from "../models/numberFormat";
import { getString } from "../models/text/text";

declare function sendCommand(c: string, data?: any): ICommandResult;

export default class BattlePage implements IPage {
  private app: IApplication;
  ractive!: Ractive;

  constructor(app: IApplication) {
    this.app = app;
  }

  async onCreate() {
    const t = await this.app.fetchTemplate("battle.html");
    this.ractive = new Ractive({
      el: "#container",
      template: t,
      components: {
        Button: Button,
        PartyStatus: PartyStatus,
      },
      data: {
        format: numberFormat,
        gs: getString,
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
