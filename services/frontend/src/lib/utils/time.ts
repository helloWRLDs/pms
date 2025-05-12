/**
 *
 * @param sec 1746759414
 * @returns 09.05.2025
 */
export const formatTime = (sec: number) => {
  const d = new Date(sec * 1000);

  const day = d.getDate().toString().padStart(2, "0");
  const month = (d.getMonth() + 1).toString().padStart(2, "0");
  const year = d.getFullYear();

  return `${day}.${month}.${year}`;
};
