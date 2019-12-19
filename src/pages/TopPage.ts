/// <reference path="./IPage.ts" />
import Ractive from "ractive";
import "../scss/top.scss";
import Button from "../views/Button";
import { handleResult } from "./commandHandler";

declare function sendCommand(c: string, data: string): ICommandResult;

export default class TopPage implements IPage {
  private app: IApplication;
  ractive!: Ractive;

  constructor(app: IApplication) {
    this.app = app;
  }

  async onCreate() {
    const t = await this.app.fetchTemplate("top.html");
    this.ractive = new Ractive({
      el: "#container",
      template: t,
      components: {
        Button: Button,
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
