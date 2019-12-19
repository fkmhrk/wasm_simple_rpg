interface IPage {
  ractive: IRactive;
  onCreate(): void;
}

interface IRactive {
  set(data: any): void;
}
