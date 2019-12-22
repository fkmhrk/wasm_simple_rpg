/// <reference path="./IApplication.ts" />
/// <reference path="./services/IServices.ts" />
/// <reference path="./clients/HTTPClient.ts" />

import { getBody, isStatus200 } from "./clients/Functions";
import TopPage from "./pages/TopPage";
import { APP_VERSION } from "./version";

declare var Go: any;
declare function init(): any;

export default class Application implements IApplication {
  private templateClient: HTTPClient;
  private templateCache: { [key: string]: string } = {};
  services: IServices;

  constructor(templateClient: HTTPClient, services: IServices) {
    this.templateClient = templateClient;
    this.services = services;
  }

  async start() {
    const go = new Go();

    const result = await WebAssembly.instantiateStreaming(
      fetch("./game.wasm"),
      go.importObject
    );
    go.run(result.instance);
    const state = init();
    console.log(state);
    new TopPage(this).onCreate();
  }

  fetchTemplate(name: string): Promise<string> {
    const template = this.templateCache[name];
    if (template != null) {
      return Promise.resolve(template);
    }
    const url = `./pages/${name}?v=${APP_VERSION}`;
    return this.templateClient
      .send(Method.GET, url, {}, null)
      .then(isStatus200)
      .then(getBody)
      .then((body: string) => {
        this.templateCache[name] = body;
        return body;
      });
  }
}
