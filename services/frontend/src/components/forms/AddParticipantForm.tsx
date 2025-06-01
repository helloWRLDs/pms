import { useState } from "react";
import { MdPersonAdd, MdEmail, MdSearch } from "react-icons/md";
import { UserFilter } from "../../lib/user/user";
import authAPI from "../../api/authAPI";
import Input from "../ui/Input";
import { Button } from "../ui/Button";

type AddParticipantFormProps = React.HTMLAttributes<HTMLDivElement> & {
  onFinish: (userID: string) => void;
  title?: string;
  placeholder?: string;
};

const AddParticipantForm = ({
  className = "",
  onFinish,
  title = "Add Participant",
  placeholder = "Enter participant's email address",
  ...props
}: AddParticipantFormProps) => {
  const [filter, setFilter] = useState<UserFilter>({
    page: 1,
    per_page: 1,
    company_id: "",
  });
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);

  const handleAddParticipant = async () => {
    if (!filter.user_email?.trim()) {
      setError("Please enter an email address");
      return;
    }

    setIsLoading(true);
    setError(null);
    setSuccess(null);

    try {
      const res = await authAPI.listUsers(filter);
      if (res && res.items && res.total_items !== 0) {
        onFinish(res.items[0].id);
        setSuccess(`Successfully added ${filter.user_email}`);
        setFilter({ ...filter, user_email: "" });
      } else {
        setError("User not found with this email address");
      }
    } catch (err) {
      console.error(err);
      setError("Failed to add participant. Please try again.");
    } finally {
      setIsLoading(false);
    }
  };

  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === "Enter") {
      handleAddParticipant();
    }
  };

  return (
    <div
      className={`bg-primary-500/20 backdrop-blur-sm rounded-xl border border-primary-400/30 p-6 shadow-lg ${className}`}
      {...props}
    >
      {/* Header */}
      <div className="flex items-center gap-3 mb-4">
        <div className="p-2 bg-accent-600 rounded-lg">
          <MdPersonAdd className="text-white" size={20} />
        </div>
        <div>
          <h3 className="text-lg font-semibold text-white">{title}</h3>
          <p className="text-white/70 text-sm">
            Search and add team members by email
          </p>
        </div>
      </div>

      {/* Form */}
      <div className="space-y-4">
        <div className="flex flex-col sm:flex-row gap-3">
          <div className="flex-1">
            <Input className="w-full">
              <Input.Element
                label=""
                type="email"
                placeholder={placeholder}
                value={filter.user_email || ""}
                onChange={(e) => {
                  setFilter({ ...filter, user_email: e.currentTarget.value });
                  setError(null);
                  setSuccess(null);
                }}
                onKeyPress={handleKeyPress}
                className="bg-white/10 border-white/20 text-white placeholder-white/50 focus:border-accent-500 focus:ring-accent-500/20"
                disabled={isLoading}
              />
            </Input>
          </div>

          <Button
            onClick={handleAddParticipant}
            disabled={isLoading || !filter.user_email?.trim()}
            className={`
              px-4 py-2 bg-accent-600 hover:bg-accent-700 text-white font-medium rounded-lg 
              transition-all duration-200 flex items-center gap-2 min-w-[80px] justify-center
              disabled:bg-white/20 disabled:cursor-not-allowed disabled:text-white/50
              hover:scale-105 hover:shadow-lg text-sm
            `}
          >
            {isLoading ? (
              <>
                <div className="w-3 h-3 border-2 border-white/30 border-t-white rounded-full animate-spin" />
                <span>Adding...</span>
              </>
            ) : (
              <>
                <MdSearch size={14} />
                <span>Add</span>
              </>
            )}
          </Button>
        </div>

        {/* Status Messages */}
        {error && (
          <div className="flex items-center gap-2 p-3 bg-red-500/20 border border-red-400/30 rounded-lg backdrop-blur-sm">
            <div className="w-2 h-2 bg-red-400 rounded-full flex-shrink-0" />
            <p className="text-red-200 text-sm">{error}</p>
          </div>
        )}

        {success && (
          <div className="flex items-center gap-2 p-3 bg-green-500/20 border border-green-400/30 rounded-lg backdrop-blur-sm">
            <div className="w-2 h-2 bg-green-400 rounded-full flex-shrink-0" />
            <p className="text-green-200 text-sm">{success}</p>
          </div>
        )}

        {/* Helper Text */}
        <div className="flex items-start gap-2 p-3 bg-blue-500/10 border border-blue-400/20 rounded-lg backdrop-blur-sm">
          <MdEmail className="text-blue-300 mt-0.5 flex-shrink-0" size={16} />
          <div className="text-blue-200 text-xs">
            <p className="font-medium mb-1">How to add participants:</p>
            <ul className="space-y-0.5 text-blue-200/80">
              <li>• Enter the exact email address of the team member</li>
              <li>• User must already be registered in the system</li>
              <li>• Press Enter or click Add to search and add</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  );
};

export default AddParticipantForm;
