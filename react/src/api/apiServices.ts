export const checkSession = async () => {
  try {
    const response = await fetch("/check-session");
    const data = await response.json();
    return data;
  } catch (error) {
    console.log("Error checking status:", error);
  }
};
