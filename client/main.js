const createNav = () => {
  const menu = [
    ["Упражнения", "exercises.html"],
    ["Добавить упражнение", "add-exercise.html"],
    ["Добавить группу мышц", "add-group.html"],
    ["Добавить пользователя", "add-user.html"],
  ];

  const navItem = (link, title) => `<li class="nav-item active">
    <a class="nav-link" href="${link}">${title}</a>
    </li>`;

  return `<nav class="navbar navbar-expand-lg navbar-light bg-light">
    <div class="container">
      <div class="navbar-collapse" id="navbarNav">
        <ul class="navbar-nav">
          ${menu.reduce((prev, curr) => prev + navItem(curr[1], curr[0]), "")}
        </ul>
      </div>
    </div>
  </nav>`;
};

document.addEventListener("DOMContentLoaded", () => {
  document.body.insertAdjacentHTML("afterbegin", createNav());
});
