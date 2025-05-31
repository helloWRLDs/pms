import { FC } from "react";
import { TreeNode } from "./TreeView";
import { useSortable } from "@dnd-kit/sortable";
import { CSS } from "@dnd-kit/utilities";
import { FiChevronDown, FiChevronRight } from "react-icons/fi";

interface DraggableTreeItemProps {
  node: TreeNode;
  level: number;
  isExpanded: boolean;
  onToggle: () => void;
  isSelected: boolean;
  onSelect: (node: TreeNode) => void;
}

export const DraggableTreeItem: FC<DraggableTreeItemProps> = ({
  node,
  level,
  isExpanded,
  onToggle,
  isSelected,
  onSelect,
}) => {
  const {
    attributes,
    listeners,
    setNodeRef,
    transform,
    transition,
    isDragging,
  } = useSortable({ id: node.id });

  const style = {
    transform: CSS.Transform.toString(transform),
    transition,
    opacity: isDragging ? 0.5 : 1,
  };

  const hasChildren = node.children && node.children.length > 0;
  const isActiveItem = node.isActive || isSelected;

  return (
    <div
      ref={setNodeRef}
      style={style}
      {...attributes}
      className={`
        group relative
        ${level > 0 ? "ml-4" : ""}
        bg-[#1a1a1a]
      `}
    >
      <div
        className={`
          flex items-center py-1 px-2 rounded-lg
          ${
            isActiveItem
              ? "bg-[#2a2a2a] text-white"
              : "text-white/70 hover:bg-[#2a2a2a]/40 hover:text-white"
          }
          ${node.className || ""}
          transition-all duration-150 ease-in-out
          cursor-pointer
          relative
          group
        `}
        onClick={() => {
          if (hasChildren) {
            onToggle();
          }
          onSelect(node);
          node.onClick?.();
        }}
        {...listeners}
      >
        <div className="flex items-center flex-1 gap-1.5">
          {hasChildren ? (
            <span className="text-white/50">
              {isExpanded ? (
                <FiChevronDown className="w-4 h-4" />
              ) : (
                <FiChevronRight className="w-4 h-4" />
              )}
            </span>
          ) : (
            <span className="w-4" /> // Spacer for alignment
          )}

          {node.icon && (
            <node.icon
              size={16}
              className={
                isActiveItem ? "text-accent-400" : "text-accent-400/70"
              }
            />
          )}

          <span className="text-sm">{node.label}</span>

          {node.badge && <span className="ml-auto">{node.badge}</span>}
        </div>
      </div>

      {hasChildren && isExpanded && (
        <div className="relative">
          <div
            className="absolute left-[7px] top-0 bottom-0 w-px bg-[#2a2a2a]"
            aria-hidden="true"
          />
          {node.children?.map((child) => (
            <DraggableTreeItem
              key={child.id}
              node={child}
              level={level + 1}
              isExpanded={false}
              onToggle={() => {}}
              isSelected={false}
              onSelect={onSelect}
            />
          ))}
        </div>
      )}
    </div>
  );
};

export default DraggableTreeItem;
