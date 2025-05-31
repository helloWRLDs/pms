export const safeNumber = (value: number | undefined | null): number => {
  if (value === undefined || value === null || isNaN(value)) {
    return 0;
  }
  return value;
};

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
