import DataFetcher from "./DataFetcher.js";
import { postsTemplate } from "./templates.js";

export default class Filters extends DataFetcher {
  static endpoint = "/api/topic/";
  static container = ".all-posts__container";
  static template = postsTemplate;

  constructor() {
    super(Filters.endpoint, Filters.container, Filters.template);
    this.topicId = this.url.substring(this.url.lastIndexOf("/") + 1);
    this.setupSelectors();
    this.addTags();
    this.handleEvents();
  }

  setupSelectors() {
    this.creationFilter = document.querySelector("#btn__byCreationDate");
    this.tagsFilter = document.querySelector("#filters__categories");
    this.tagsButton = document.querySelector("#btn__byCategory");
    this.clearFilters = document.querySelector("#clear__filters");
    this.filtersContainer = document.querySelector(
      "#creationDate__filters__container"
    );
    this.sortByLikesBtn = document.querySelector("#btn__mostLikes");
    this.oldestFirst = document.querySelector("#creationDate__filters__oldest");
    this.newestFirst = document.querySelector("#creationDate__filters__newest");
    this.uniqueElements = [];
    this.filteredTagsData = [];
    this.checkboxes = null;
    this.isOpened = false;
  }

  handleEvents() {
    this.sortByLikesBtn.addEventListener(
      "click",
      this.toggleSort.bind(this, "likes")
    );
    this.creationFilter.addEventListener("mouseenter", () =>
      this.filtersContainer.classList.remove("hidden")
    );
    this.creationFilter.addEventListener("mouseleave", () =>
      this.filtersContainer.classList.add("hidden")
    );
    this.oldestFirst.addEventListener(
      "click",
      this.toggleSort.bind(this, "oldest")
    );
    this.newestFirst.addEventListener(
      "click",
      this.toggleSort.bind(this, "newest")
    );
    this.clearFilters.addEventListener("click", () => window.location.reload());
  }

  toggleSort(type) {
    if (type === "likes") {
      this.sortByLikes();
      this.creationFilter.classList.remove("active-bg");
      this.sortByLikesBtn.classList.add("active-bg");
    } else if (type === "oldest" || type === "newest") {
      this.sortByLikesBtn.classList.remove("active-bg");
      this.newestFirst.classList.toggle("active-filter", type === "newest");
      this.oldestFirst.classList.toggle("active-filter", type === "oldest");
      this.creationFilter.classList.add("active-bg");
      this.sortByDate(type);
    }
  }

  sortByLikes() {
    this.elements = document.querySelectorAll(".all-posts__item");
    this.oldestFirst.classList.remove("active-filter");
    this.newestFirst.classList.remove("active-filter");
    const sortedElements = Array.from(this.elements).sort((a, b) => {
      const likesA = parseInt(a.dataset.likes);
      const likesB = parseInt(b.dataset.likes);
      return likesB - likesA;
    });

    sortedElements.forEach((el) => {
      this.container.appendChild(el);
    });
  }

  sortByDate(order) {
    this.elements = document.querySelectorAll(".all-posts__item");

    const sortedElements = Array.from(this.elements).sort((a, b) => {
      const dateA = new Date(a.dataset.created);
      const dateB = new Date(b.dataset.created);
      return order === "oldest" ? dateA - dateB : dateB - dateA;
    });

    sortedElements.forEach((el) => {
      this.container.appendChild(el);
    });
  }

  sortByTags() {
    this.elements = document.querySelectorAll(".all-posts__item");

    const checkboxes = Array.from(
      document.querySelectorAll(".filter-checkbox")
    );
    const selectedTags = checkboxes
      .filter((cb) => cb.checked)
      .map((cb) => cb.id);

    if (selectedTags.length === 0) {
      this.elements.forEach((el) => (el.style.display = ""));
      return;
    }

    this.elements.forEach((post) => {
      const postTags = post.dataset.tags.split(",");
      const isMatch = selectedTags.some((selectedTag) =>
        postTags.includes(selectedTag)
      );
      post.style.display = isMatch ? "" : "none";
    });
  }

  async addTags() {
    try {
      const response = await fetch(`${Filters.endpoint}${this.topicId}`);
      const data = await response.json();
      this.originalData = data;
      const allTags = data.map((item) => item.tags.split(","));
      allTags.flat().forEach((tag) => {
        if (!this.uniqueElements.includes(tag)) {
          this.uniqueElements.push(tag);
        }
      });
      this.tagsFilter.innerHTML = "";
      this.uniqueElements.forEach((element) => {
        this.tagsFilter.insertAdjacentHTML(
          "afterbegin",
          ` 
            <li class="filters__categories__item">
              <input
                class="filter-checkbox"
                type="checkbox"
                name="check-1"
                value="${element}"
                id="${element}"
              />
              <label for="${element}">#${element}</label>
            </li>
          `
        );
        const checkbox = document.getElementById(element);
        checkbox.addEventListener("change", () => {
          this.sortByTags();
        });
      });
    } catch (error) {
      console.error("Error fetching data: ", error);
    }
  }
}
