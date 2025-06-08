/**
 * Returns the safe number of a value
 * @param value - The value to get the safe number of
 * @returns The safe number of the value
 */
export const safeNumber = (value: number | undefined | null): number => {
  if (value === undefined || value === null || isNaN(value)) {
    return 0;
  }
  return value;
};

/**
 * Returns the safe percentage of a numerator and denominator
 * @param numerator - The numerator of the percentage
 * @param denominator - The denominator of the percentage
 * @returns The safe percentage of the numerator and denominator
 */
export const safePercentage = (
  numerator: number,
  denominator: number
): string => {
  const num = safeNumber(numerator);
  const den = safeNumber(denominator);

  if (den === 0) {
    return "0";
  }

  const percentage = (num / den) * 100;

  if (isNaN(percentage) || !isFinite(percentage)) {
    return "0";
  }

  return percentage.toFixed(0);
};

export const safeDisplay = (value: number | undefined | null): string => {
  const num = safeNumber(value);
  return isNaN(num) ? "0" : num.toString();
};
