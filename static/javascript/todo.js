      const input = document.getElementById("input");
      const formSubmit = document.getElementById("submitForm");
      const listItem = document.querySelector(".ul-list");
      const circle = document.querySelectorAll(".circle");
      const url = "http://localhost:8080/api/todo";

      function start() {
        getAPi(renderHTML);
        postItem();
      }
      function getAPi(callback) {
        fetch(url)
          .then((response) => response.json())
          .then(callback);
      }
      function renderHTML(datas) {
        input.value = "";
        input.focus();
        var i = -1;
        var htmls = datas.map((data, id) => {
          i++;
          return data.status
            ? `<li class="list">
              <div class="circle" onclick="handleCheck(${i},event)">
                <i class="fa-solid fa-check"></i>
              </div>
              <div class="txt"
                  ondblclick="handleDoubleEdit(${i},event)">
                ${data.title}
                <input
                  type="text"
                  id="formEdit"
                  class="formEdit block"
                  placeholder="..."
                  autocomplete="off"
                  maxlength="47"
                />
              </div>
              <div class="option">
                <span class="linkEdit" onclick="handleEdit(${i},event)"
                  ><div class="btn editItem">
                    <i class="fa-regular fa-pen-to-square"></i></div
                ></span>
                <span class="linkDel" onclick="handledelete(${i})"
                  ><div class="btn deleItem">
                    <i class="fa-solid fa-trash-can"></i></div
                ></span>
              </div>
            </li>`
            : `<li class="list">
              <div class="circle active" onclick="handleCheck(${i},event)">
                <i class="fa-solid fa-check"></i>
              </div>
              <div class="txt"
                  ondblclick="handleDoubleEdit(${i},event)">
                ${data.title}
                <input
                  type="text"
                  id="formEdit"
                  class="formEdit block"
                  placeholder="..."
                  autocomplete="off"
                  maxlength="47"
                />
              </div>
              <div class="option active">
                <span class="linkEdit" onclick="handleEdit(${i},event)"
                  ><div class="btn editItem">
                    <i class="fa-regular fa-pen-to-square"></i></div
                ></span>
                <span class="linkDel" onclick="handledelete(${i})"
                  ><div class="btn deleItem">
                    <i class="fa-solid fa-trash-can"></i></div
                ></span>
              </div>
            </li>`;
        });
        listItem.innerHTML = htmls.join("");
      }

      function postItem() {
        if (formSubmit.getAttribute("onsubmit") == "true") {
          const exr = input.value;
          const options = {
            method: "POST",
            body: JSON.stringify({ title: exr, status: true }),
          };
          fetch(url, options)
            .then((response) => response.json())
            .catch((error) => error.json());
          getAPi(renderHTML);
        } else {
          return;
        }
      }
      function handleEdit(id, event) {
        event.target
          .closest(".list")
          .querySelector(".formEdit")
          .classList.toggle("block");
        const context = event.target
          .closest(".list")
          .querySelector(".formEdit");
        context.addEventListener("keypress", (even) => {
          if (even.key == "Enter") {
            if (context.value == "") {
              return;
            } else {
              const options = {
                method: "PUT",
                body: JSON.stringify({ title: context.value, status: true }),
              };
              fetch(url + "/" + id, options)
                .then((response) => response.json())
                .catch((error) => error.json());
              event.target
                .closest(".list")
                .querySelector(".formEdit")
                .classList.toggle("block");
              getAPi(renderHTML);
              location.reload();
            }
            return;
          }
        });
      }
      function handledelete(id) {
        const options = {
          method: "DELETE",
        };
        fetch(url + "/" + id, options)
          .then((response) => response.json())
          .catch((error) => error.json());
        getAPi(renderHTML);
        location.reload();
      }

      start();

      formSubmit.addEventListener("submit", function (e) {
        e.preventDefault();
      });

      function handleCheck(id, event) {
        const context = event.target.closest(".list").querySelector(".txt");
        var bool;
        event.target.closest(".circle").classList.toggle("active")
          ? (bool = false)
          : (bool = true);
        const options = {
          method: "PUT",
          body: JSON.stringify({
            title: context.innerText,
            status: bool,
          }),
        };
        fetch(url + "/" + id, options)
          .then((response) => response.json())
          .catch((error) => error.json());
      }
      function handleDoubleEdit(id) {
        handleEdit(id, event);
      }
      input.addEventListener("keypress", (even) => {
        if (even.key == "Enter") {
          if (input.value == "") {
            formSubmit.setAttribute("onsubmit", false);
            start();
          } else {
            formSubmit.setAttribute("onsubmit", true);
            start();
          }
          return;
        }
      });