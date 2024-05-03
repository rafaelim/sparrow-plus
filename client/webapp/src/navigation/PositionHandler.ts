type Sections = "movies" | "watch-next" | "tv-shows" | "episodes";

class PositionHandler {
  private Y: number;
  private sections: Map<Sections, number>;

  constructor() {
    this.sections = new Map();
    this.Y = 0;
  }

  getNextY(section: Sections) {
    if (this.sections.has(section)) return this.sections.get(section) as number;
    const X = this.Y;
    this.sections.set(section, X);
    this.Y += 1;
    return X;
  }

  getLastY() {
    return this.Y - 1;
  }

  addListener(listener: () => void) {
    window.addEventListener("open-item", listener);

    return () => window.removeEventListener("open-item", listener);
  }

  triggerOpenEvent() {
    const evt = new Event("open-item");
    window.dispatchEvent(evt);
  }
}

export default new PositionHandler();
