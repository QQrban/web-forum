import { hideLoader, showLoader } from "./helpers.js";

export default class DataFetcher {
  constructor(endpoint, container, template) {
    this.container = document.querySelector(container);
    this.endpoint = endpoint;
    this.url = window.location.pathname;
    this.template = template;
    this.sortedData = [];
    this.topicId = this.url.substring(this.url.lastIndexOf("/") + 1);
    showLoader();
  }

  async fetchByID(order) {
    try {
      const response = await fetch(`${this.endpoint}${this.topicId}`);
      const data = await response.json();
      this.container.innerHTML = "";
      const newData = this.filterByDate(data, "created_raw", order);
      this.displayContent(newData);
    } catch (error) {
      console.error("Error fetching posts:", error);
    } finally {
      setTimeout(() => {
        hideLoader();
      }, 200);
    }
  }

  filterByDate(data, filterBy, order) {
    this.sortedData = [...data].sort((a, b) => {
      return order === "descending"
        ? new Date(b[filterBy]) - new Date(a[filterBy])
        : new Date(a[filterBy]) - new Date(b[filterBy]);
    });
    return this.sortedData.length > 0 ? this.sortedData : data;
  }

  displayContent(data) {
    let i = 0;
    data.forEach((dataItem) => {
      i++;
      let template = this.template
        .replace(/{{title}}/g, dataItem.title)
        .replace(/{{author}}/g, dataItem.author)
        .replace(/{{avatar}}/g, dataItem.avatar)
        .replace(/{{created}}/g, dataItem.created)
        .replace(/{{created_raw}}/g, dataItem.created_raw)
        .replace(/{{likes}}/g, dataItem.likes)
        .replace(/{{post_id}}/g, dataItem.post_id)
        .replace(/{{dislikes}}/g, dataItem.dislikes)
        .replace(/{{comments_number}}/g, dataItem.comments_number)
        .replace(/{{last_commentator_name}}/g, dataItem.last_commentator_name)
        .replace(/{{content}}/g, dataItem.content)
        .replace(/{{posts_number}}/g, dataItem.posts_number)
        .replace(/{{member_days}}/g, dataItem.member_days)
        .replace(/{{comment_id}}/g, dataItem.comment_id)
        .replace(/{{tags}}/g, dataItem.tags)
        .replace(/{{count}}/g, i)
        .replace(
          /{{last_commentator_avatar}}/g,
          dataItem.last_commentator_avatar
        )
        .replace(/{{last_comment_created}}/g, dataItem.last_comment_created);
      this.container.insertAdjacentHTML("beforeend", template);
    });
  }
}
