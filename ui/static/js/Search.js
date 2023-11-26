export default class Search {
  constructor() {
    this.searchInput = document.querySelector(".search-bar__input");
    this.resultContainer = document.querySelector(".search-bar__results");
    this.addEventListener();
  }

  addEventListener() {
    this.searchInput.addEventListener("keyup", () => this.handleSearch());
    document.addEventListener("click", (e) => {
      if (!e.target.classList.contains("search-bar__input")) {
        this.resultContainer.classList.remove("active-search");
      }
    });
  }

  async handleSearch() {
    if (this.searchInput.value.length >= 2) {
      this.handleClick();
      this.resultContainer.classList.add("active-search");
      const response = await fetch("/api/all-posts");
      if (!response.ok) {
        throw new Error("HTTP Status Error:", response.statusText);
      } else {
        const data = await response.json();
        if (!Array.isArray(data)) {
          this.resultContainer.innerHTML = "";
          return;
        }
        this.resultContainer.innerHTML = "";
        for (const suggestion of data) {
          if (
            suggestion.title
              .toLowerCase()
              .includes(this.searchInput.value.toLowerCase())
          ) {
            this.resultContainer.insertAdjacentHTML(
              "afterbegin",
              `
              <a href="/post/${suggestion.post_id}">
                <li class="search-bar__result-item">
                    ${suggestion.title}
                </li>
              </a>
            `
            );
          }
        }
      }
    } else {
      this.resultContainer.classList.remove("active-search");
    }
  }

  handleClick() {
    this.searchInput.addEventListener("click", () => {
      if (this.searchInput.value.length >= 2) {
        this.resultContainer.classList.add("active-search");
      }
    });
  }
}
