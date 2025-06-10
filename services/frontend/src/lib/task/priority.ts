import { ObjectWrapper } from "../utils/wrapper";

export enum PriorityLevel {
  ALL = 0,
  LOWEST = 1,
  LOW = 2,
  MEDIUM = 3,
  HIGH = 4,
  CRITICAL = 5,
}

export type PriorityConfig = {
  label: string;
  color: string;
};

const PRIORITY_CONFIG: Record<PriorityLevel, PriorityConfig> = {
  [PriorityLevel.ALL]: {
    label: "All",
    color: "gray",
  },
  [PriorityLevel.LOWEST]: {
    label: "Lowest",
    color: "green",
  },
  [PriorityLevel.LOW]: {
    label: "Low",
    color: "yellow",
  },
  [PriorityLevel.MEDIUM]: {
    label: "Medium",
    color: "orange",
  },
  [PriorityLevel.HIGH]: {
    label: "High",
    color: "red",
  },
  [PriorityLevel.CRITICAL]: {
    label: "Critical",
    color: "purple",
  },
};

export const AVAILABLE_PRIORITIES = [
  PriorityLevel.LOWEST,
  PriorityLevel.LOW,
  PriorityLevel.MEDIUM,
  PriorityLevel.HIGH,
  PriorityLevel.CRITICAL,
];

export class Priority {
  constructor(private readonly level: PriorityLevel) {}

  static getAvailablePriorities(): PriorityLevel[] {
    return AVAILABLE_PRIORITIES;
  }

  toString(): string {
    return PRIORITY_CONFIG[this.level].label;
  }

  getColor(): string {
    return PRIORITY_CONFIG[this.level].color;
  }

  getLevel(): PriorityLevel {
    return this.level;
  }

  static fromNumber(value: number): Priority {
    if (!(value in PriorityLevel)) {
      return new Priority(PriorityLevel.ALL);
    }
    return new Priority(value as PriorityLevel);
  }
}

export const PriorityFilterValues: ObjectWrapper[] = [
  PriorityLevel.ALL,
  ...AVAILABLE_PRIORITIES,
].map((level) => ({
  label: PRIORITY_CONFIG[level].label,
  value: level,
}));
