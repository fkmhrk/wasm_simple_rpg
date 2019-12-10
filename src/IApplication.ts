/// <reference path="./services/IServices.ts" />

interface IApplication {
  services: IServices;

  start(): void;

  fetchTemplate(name: string): Promise<string>;
}
