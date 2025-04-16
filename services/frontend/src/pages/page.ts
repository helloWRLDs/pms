export class PageSettings {
  title: string;
  showHeader: boolean;
  showFooter: boolean;
  showSidebar: boolean;

  constructor(
    title: string,
    showHeader: boolean = true,
    showFooter: boolean = true,
    showSidebar: boolean = true
  ) {
    this.title = title;
    this.showHeader = showHeader;
    this.showFooter = showFooter;
    this.showSidebar = showSidebar;
  }

  setup() {
    document.title = this.title;
  }
}
