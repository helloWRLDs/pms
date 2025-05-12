import { PropsWithChildren } from "react";
import { IoCloseSharp } from "react-icons/io5";
import { useOutsideClick } from "../../hooks/useModal";

interface Props extends PropsWithChildren {
  onClose: () => void;
  visible: boolean;
  title: string;
  className?: string;
}

export const Modal = ({ visible = false, ...props }: Props) => {
  const modalRef = useOutsideClick(props.onClose);
  if (!visible) {
    return null;
  }
  return (
    <div
      className={`${props.className} px-10 py-5 rounded-md shadow-2xl shadow-gray-500`}
      ref={modalRef}
    >
      <div className="modal-header">
        <button className="modal-close-btn" onClick={props.onClose}>
          <IoCloseSharp size="30px" color="white" />
        </button>
        <h3 className="modal-title">{props.title ?? ""}</h3>
      </div>
      {props.children ?? null}
    </div>
  );
};
