import { FC, ReactNode, useState } from "react";
import { IconType } from "react-icons";
import { FiChevronRight, FiChevronDown } from "react-icons/fi";
// import { VscCircleFilled } from "react-icons/vsc";

export interface TreeNode {
  id: string;
  label: string;
  icon?: IconType;
  isActive?: boolean;
  onClick?: () => void;
  children?: TreeNode[];
  className?: string;
  badge?: ReactNode;
}

interface TreeViewProps {
  nodes: TreeNode[];
  defaultExpandedIds?: string[];
  activeNodeId?: string;
  onNodeSelect?: (node: TreeNode) => void;
  className?: string;
}

interface TreeItemProps {
  node: TreeNode;
  level: number;
  isExpanded: boolean;
  onToggle: () => void;
  isSelected: boolean;
  onSelect: (node: TreeNode) => void;
}

const TreeItem: FC<TreeItemProps> = ({
  node,
  level,
  isExpanded,
  onToggle,
  isSelected,
  onSelect,
}) => {
  const hasChildren = node.children && node.children.length > 0;
  const isActiveItem = node.isActive || isSelected;

  return (
    <div
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
            <TreeNodeContainer
              key={child.id}
              node={child}
              level={level + 1}
              onSelect={onSelect}
              expandedNodeId={null}
              onNodeExpand={() => {}}
            />
          ))}
        </div>
      )}
    </div>
  );
};

const TreeNodeContainer: FC<{
  node: TreeNode;
  level: number;
  onSelect: (node: TreeNode) => void;
  expandedNodeId: string | null;
  onNodeExpand: (nodeId: string) => void;
}> = ({ node, level, onSelect, expandedNodeId, onNodeExpand }) => {
  const isSelected = node.isActive ?? false;
  const isExpanded =
    level === 0 ? node.id === expandedNodeId : Boolean(node.isActive);
  const [isSubTreeExpanded, setIsSubTreeExpanded] = useState(true);

  const handleToggle = () => {
    if (level === 0) {
      onNodeExpand(node.id);
    } else {
      setIsSubTreeExpanded(!isSubTreeExpanded);
    }
  };

  return (
    <TreeItem
      node={node}
      level={level}
      isExpanded={level === 0 ? isExpanded : isSubTreeExpanded}
      onToggle={handleToggle}
      isSelected={isSelected}
      onSelect={onSelect}
    />
  );
};

export const TreeView: FC<TreeViewProps> = ({
  nodes,
  className = "",
  onNodeSelect,
}) => {
  const [expandedNodeId, setExpandedNodeId] = useState<string | null>(null);

  const handleNodeExpand = (nodeId: string) => {
    setExpandedNodeId(expandedNodeId === nodeId ? null : nodeId);
  };

  const handleSelect = (node: TreeNode) => {
    onNodeSelect?.(node);
  };

  return (
    <div className={`py-2 ${className} bg-[#1a1a1a]`}>
      {nodes.map((node) => (
        <TreeNodeContainer
          key={node.id}
          node={node}
          level={0}
          onSelect={handleSelect}
          expandedNodeId={expandedNodeId}
          onNodeExpand={handleNodeExpand}
        />
      ))}
    </div>
  );
};

export default TreeView;
