import { FC, useEffect, useState } from "react";
import DialogModal from "../components/ui/DialogModal";
import { PageSettings } from "./page";

class ProfilePageSettings extends PageSettings {}

const ProfilePage: FC = () => {
  const settings = new ProfilePageSettings("Profile");
  const [open, setOpen] = useState(true);

  useEffect(() => {
    settings.setup();
  }, []);

  return (
    <div className="flex justify-between bg-primary-200">
      <h1>Profile</h1>
      <DialogModal
        open={open}
        onClose={() => setOpen(false)}
        title="Warning!"
        type="error"
        confirmText="Proceed"
        cancelText="Cancel"
        onConfirm={() => console.log("User confirmed")}
      >
        <p>This action might have serious consequences. Are you sure?</p>
      </DialogModal>
    </div>
  );
};

export default ProfilePage;
