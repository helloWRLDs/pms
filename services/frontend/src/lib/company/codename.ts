const MAX_NON_VOWELS = 5;
const REGEX = /(?<=\p{Ll})(?=\p{Lu})|(?<=\D)(?=\d)|(?<=\d)(?=\D)|\s+/gu;
const SKIP_KEY_WORDS = [
  "LLC",
  "LLP",
  "PVT",
  "LTD",
  "PVT LTD",
  "PVT LTD",
  "ТОО",
  "ИП",
  "ООО",
  "ОАО",
  "ЗАО",
  "НПО",
  "НПО",
  "НПО",
];
const VOWELS = [
  "а",
  "е",
  "и",
  "о",
  "у",
  "ы",
  "э",
  "ю",
  "я",
  "a",
  "e",
  "i",
  "o",
  "u",
  "y",
];

export const generateCodeName = (name: string) => {
  const rawParts = name.split(REGEX);
  const parts = rawParts.filter((part) => !SKIP_KEY_WORDS.includes(part));
  if (parts.length === 1) {
    const core = parts[0].toLowerCase();
    let result = "";

    for (const char of core) {
      if (!VOWELS.includes(char.toLowerCase()) && /[a-zа-яё]/i.test(char)) {
        result += char;
        if (result.length >= MAX_NON_VOWELS) break;
      }
    }

    return result;
  }
  let result = "";
  for (let part of parts) {
    const lower = part.toLowerCase();

    if (!isNaN(Number(part))) {
      result += part;
    } else if (part.length < 4) {
      result += lower;
    } else {
      result += lower[0];
    }
  }
  return result;
};
