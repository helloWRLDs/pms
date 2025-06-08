import { FC, useEffect } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";
import { useAuthStore } from "../../store/authStore";
import { toast } from "react-toastify";
import { toastOpts } from "../../lib/utils/toast";

// Helper function to properly decode base64 with Unicode support
function decodeBase64(str: string): string {
  // First, we need to handle the base64 padding
  str = str.replace(/-/g, "+").replace(/_/g, "/");
  // Add padding if needed
  while (str.length % 4) str += "=";

  // Use TextDecoder to properly handle Unicode
  const binary = atob(str);
  const bytes = new Uint8Array(binary.length);
  for (let i = 0; i < binary.length; i++) {
    bytes[i] = binary.charCodeAt(i);
  }
  return new TextDecoder().decode(bytes);
}

const OAuthCallback: FC = () => {
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();
  const { setAuth } = useAuthStore();

  useEffect(() => {
    const handleOAuthCallback = async () => {
      console.log("🔄 Processing OAuth callback...");

      // Check for error from backend
      const error = searchParams.get("error");
      if (error) {
        console.error("❌ OAuth error from backend:", error);
        toast.error(`OAuth error: ${error}`, toastOpts);
        navigate("/login");
        return;
      }

      // Check for success flag and auth data
      const success = searchParams.get("success");
      const authData = searchParams.get("data");

      if (success === "true" && authData) {
        try {
          console.log("📦 Processing auth data from backend...");

          // Decode base64 auth data with proper Unicode handling
          const decodedData = decodeBase64(decodeURIComponent(authData));
          console.log("Decoded auth data:", decodedData);
          const authResponse = JSON.parse(decodedData);

          console.log("Auth response decoded:", JSON.stringify(authResponse));

          if (authResponse.payload) {
            // Set authentication data in store
            setAuth(authResponse.payload);
            toast.success("Successfully signed in with Google!", toastOpts);
            console.log("🎉 Authentication successful, redirecting to home...");
            navigate("/");
          } else {
            throw new Error("Invalid auth response structure");
          }
        } catch (parseError) {
          console.error("❌ Failed to parse auth data:", parseError);
          toast.error("Failed to process authentication data", toastOpts);
          navigate("/login");
        }
        return;
      }

      // Fallback: if no success flag but no error, something went wrong
      console.error("❌ Invalid OAuth callback - no success flag or auth data");
      toast.error("Invalid OAuth callback parameters", toastOpts);
      navigate("/login");
    };

    handleOAuthCallback();
  }, [searchParams, navigate, setAuth]);

  return (
    <div className="flex justify-center items-center min-h-screen bg-gradient-to-br from-primary-600 via-primary-700 to-primary-800">
      <div className="bg-white p-8 rounded-xl shadow-2xl">
        <div className="text-center">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-accent-500 mx-auto mb-4"></div>
          <h2 className="text-xl font-semibold text-gray-900 mb-2">
            Completing sign-in...
          </h2>
          <p className="text-gray-600">
            Please wait while we verify your account.
          </p>
        </div>
      </div>
    </div>
  );
};

export default OAuthCallback;
