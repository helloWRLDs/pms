export const fileToBase64 = (file: File): Promise<string> => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => {
      const result = reader.result as string;
      // Remove the data:image/jpeg;base64, prefix
      resolve(result.split(",")[1]);
    };
    reader.onerror = reject;
  });
};

// Validate image file
export const validateImageFile = (
  file: File
): { isValid: boolean; error?: string } => {
  // Validate file type
  if (!file.type.startsWith("image/")) {
    return { isValid: false, error: "Please select a valid image file" };
  }

  // Validate file size (max 5MB)
  if (file.size > 5 * 1024 * 1024) {
    return { isValid: false, error: "Image size must be less than 5MB" };
  }

  return { isValid: true };
};

// Validate URL
export const validateUrl = (
  url: string
): { isValid: boolean; error?: string } => {
  if (!url.trim()) {
    return { isValid: false, error: "Please enter a valid URL" };
  }

  try {
    new URL(url);
    return { isValid: true };
  } catch {
    return { isValid: false, error: "Please enter a valid URL" };
  }
};
