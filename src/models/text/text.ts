const jaStrings: { [key: string]: string } = {
  battle_found: "{}に、みつかった！",
  battle_option: "{}は、どうする？",
  battle_damage: "{}に、{}のダメージ！",
  battle_killed: "{}は、やられてしまった！",

  next_level: "{}は、レベル{}に、なった！",
};
const defaultStrings: { [key: string]: string } = {};

const strings: { [key: string]: { [key: string]: string } } = {
  "ja-JP": jaStrings,
};

const lang = window.navigator.language;

export const getString = (key: string, arg1?: any, arg2?: any) => {
  let strMap = strings[lang];
  if (strMap == null) {
    strMap = defaultStrings;
  }
  let str = strMap[key];
  if (str == null) {
    return "<undefined>";
  }
  if (arg1 !== undefined) {
    str = str.replace("{}", arg1);
  }
  if (arg2 !== undefined) {
    str = str.replace("{}", arg2);
  }
  return str;
};
