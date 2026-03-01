const createNav = () => {
  const menu = [
    ["Упражнения", "exercises.html"],
    ["Добавить упражнение", "add-exercise.html"],
    ["Добавить группу мышц", "add-group.html"],
    ["Добавить пользователя", "add-user.html"],
  ];

  const navItem = (link, title) =>
    `<li class="nav-item active">
    <a class="nav-link" href="${link}">${title}</a>
  </li>`;

  const navStyle = `<style>
    a.nav-link {
      color: blue;
      text-decoration: underline;
    }
    a.nav-link:hover {
      color: blue;
      text-decoration: none;
    }
  </style>`;
  const navHtml = `<nav class="navbar navbar-expand-lg navbar-light bg-light">
    <div class="container">
      <div class="navbar-collapse" id="navbarNav">
        <ul class="navbar-nav">
          ${menu.reduce((prev, curr) => prev + navItem(curr[1], curr[0]), "")}
        </ul>
      </div>
    </div>
  </nav>`;

  return navStyle + "\n" + navHtml;
};



const showResponse = (message) => {
  const $form = document.querySelector("form");
  const $formButton = document.querySelector("form button");

  if (!$form || !$formButton) return;

  const existing = $form.querySelector(".alert.alert-primary");
  if (existing) existing.remove();

  
  const wrap = document.createElement("div");
  wrap.className = "alert alert-primary";
  wrap.setAttribute("role", "alert");
  wrap.textContent = message;
  $formButton.insertAdjacentElement('beforebegin', wrap)
};

function showError(message) {
  const $form = document.querySelector("form");
  const $formButton = document.querySelector("form button");
  if (!$form || !$formButton) return;
  if (message) {
    const existing = $form.querySelector(".alert.alert-danger");
    if (existing) existing.remove();
    const wrap = document.createElement("div");
    wrap.className = "alert alert-danger";
    wrap.setAttribute("role", "alert");
    wrap.textContent = message;
    $formButton.insertAdjacentElement('beforebegin', wrap)
  } else {
    const alert = $form.querySelector(".alert.alert-danger");
    if (alert) alert.remove();
  }
}

function apiRequest(url, config = {}) {
  const { method = "GET", token: needToken = false, body } = config;
  let onSuccessFn = () => {};
  let onErrorFn = () => {};

  const run = async () => {
    const headers = {};
    if (needToken) {
      headers.authorization = sessionStorage.getItem("token") || "";
    }
    let requestBody = body;
    if (body && typeof body === "object" && !(body instanceof FormData)) {
      headers["Content-Type"] = "application/json";
      requestBody = JSON.stringify(body);
    }
    try {
      showError();
      showResponse();
      const res = await fetch(url, { method, headers, body: requestBody });
      if (!res.ok) {
        const text = await res.text();
        const err = { message: text || "Ошибка сервера" };
        showError(err.message);
        onErrorFn(err);
        return;
      }
      const data = await res.json().catch(() => ({}));
      onSuccessFn(data);
    } catch (err) {
      const error = { message: err?.message || "Ошибка сервера" };
      showError(error.message);
      onErrorFn(error);
    }
  };

  return {
    onSuccess(fn) {
      onSuccessFn = fn;
      return this;
    },
    onError(fn) {
      onErrorFn = fn;
      return this;
    },
    run,
  };
}

document.addEventListener("DOMContentLoaded", () => {
  document.body.insertAdjacentHTML("afterbegin", createNav());
});
