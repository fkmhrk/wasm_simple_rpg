import Ractive from "ractive";
import TopPage from "./TopPage";
import MovePage from "./MovePage";

export const handleResult = async (
  result: ICommandResult,
  app: IApplication,
  r: Ractive
) => {
  if (result.error_code != null) {
    console.log(result.error_code);
    return;
  }
  if (result.next_page.length == 0) {
    r.set(result.data);
    return;
  }

  let nextPage: IPage | null = null;
  switch (result.next_page) {
    case "top":
      nextPage = new TopPage(app);
      break;
    case "move":
      nextPage = new MovePage(app);
      break;
  }
  if (nextPage != null) {
    await nextPage.onCreate();
    nextPage.ractive.set(result.data);
  }
};
