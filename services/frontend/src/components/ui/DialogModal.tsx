import { FC, ReactNode } from "react";
import {
  Dialog,
  DialogBackdrop,
  DialogPanel,
  DialogTitle,
  Transition,
  TransitionChild,
} from "@headlessui/react";
import {
  FaExclamationTriangle,
  FaCheckCircle,
  FaInfoCircle,
  FaTimesCircle,
} from "react-icons/fa";

interface DialogModalProps {
  open: boolean;
  onClose: () => void;
  title: string;
  type?: "info" | "warn" | "success" | "error";
  children: ReactNode;
  confirmText?: string;
  cancelText?: string;
  onConfirm?: () => void;
}

const ICONS = {
  info: { icon: FaInfoCircle, color: "text-blue-400 bg-blue-100" },
  warn: { icon: FaExclamationTriangle, color: "text-yellow-600 bg-yellow-100" },
  success: { icon: FaCheckCircle, color: "text-green-600 bg-green-100" },
  error: { icon: FaTimesCircle, color: "text-red-600 bg-red-100" },
};

const DialogModal: FC<DialogModalProps> = ({
  open,
  onClose,
  title,
  type = "info",
  children,
  confirmText = "Confirm",
  cancelText = "Cancel",
  onConfirm,
}) => {
  const { icon: Icon, color } = ICONS[type];

  return (
    <Transition show={open}>
      <Dialog onClose={onClose} className="relative z-50">
        {/* Backdrop with fade */}
        <TransitionChild
          enter="ease-out duration-300"
          enterFrom="opacity-0"
          enterTo="opacity-100"
          leave="ease-in duration-200"
          leaveFrom="opacity-100"
          leaveTo="opacity-0"
        >
          <DialogBackdrop className="fixed inset-0 bg-black/30" />
        </TransitionChild>

        <div className="fixed inset-0 z-50 flex items-center justify-center p-4">
          {/* Dialog Panel with scale/opacity animation */}
          <TransitionChild
            enter="ease-out duration-300"
            enterFrom="opacity-0 scale-95"
            enterTo="opacity-100 scale-100"
            leave="ease-in duration-200"
            leaveFrom="opacity-100 scale-100"
            leaveTo="opacity-0 scale-95"
          >
            <DialogPanel className="w-full max-w-md transform  overflow-hidden rounded-xl bg-primary-100 p-6 shadow-xl transition-all">
              <div className="flex items-center space-x-4 py-2">
                <div
                  className={`flex bg-transparent items-center justify-center rounded-lg ${color} `}
                >
                  <Icon className="size-6" />
                </div>
                <DialogTitle
                  as="h3"
                  className="text-lg font-semibold text-soft-500"
                >
                  {title}
                </DialogTitle>
              </div>

              <div className="mt-4 text-soft-600">{children}</div>

              <div className="mt-6 flex justify-end gap-3">
                <button
                  type="button"
                  onClick={onClose}
                  className="cursor-pointer inline-flex justify-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50"
                >
                  {cancelText}
                </button>
                <button
                  type="button"
                  onClick={() => {
                    onConfirm?.();
                    onClose();
                  }}
                  className={`cursor-pointer inline-flex justify-center rounded-md px-4 py-2 text-sm font-medium text-white shadow-sm ${
                    type === "success"
                      ? "bg-green-600 hover:bg-green-500"
                      : type === "warn"
                      ? "bg-yellow-600 hover:bg-yellow-500"
                      : type === "error"
                      ? "bg-red-600 hover:bg-red-500"
                      : "bg-blue-600 hover:bg-blue-500"
                  }`}
                >
                  {confirmText}
                </button>
              </div>
            </DialogPanel>
          </TransitionChild>
        </div>
      </Dialog>
    </Transition>
  );
};

export default DialogModal;
