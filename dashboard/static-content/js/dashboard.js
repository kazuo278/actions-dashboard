// 現在のページ数
let currentPageNum = 1;
// 1ページあたりの表示数
let pageSize = 5;

// 検索URLを組み立てる関数
let createRequestUri = () => {
  var uri = window.location.protocol + "//" + window.location.host + "/actions/history";
  var params = new URLSearchParams();

  params.append("limit", pageSize);
  params.append("offset", (currentPageNum - 1) * pageSize);

  var repository_id = document.getElementById("repository_id").value;
  if(repository_id){
    params.append("repository_id", repository_id);
  }

  var repository_name = document.getElementById("repository_name").value;
  if(repository_name){
    params.append("repository_name", repository_name);
  }

  var started_at = document.getElementById("started_at").value;
  if (started_at) {
    params.append("started_at", started_at + "T00:00:00Z");
  }

  var finished_at = document.getElementById("finished_at").value;
  if (finished_at) {
    params.append("finished_at", finished_at + "T00:00:00Z");
  }

  var status = document.getElementById("status").value
  if (status !== "ALL"){
    params.append("status", status);
  }
  uri += "?" + new URLSearchParams(params).toString();
  return uri;
}

// 日付変換する関数
let formatDate = (dateStr) => {
  date = new Date(dateStr);
  var year = date.getFullYear().toString().padStart(4, "0");
  var month = (date.getMonth() + 1).toString().padStart(2, "0");
  var day = date.getDate().toString().padStart(2, "0");
  var hours = date.getHours().toString().padStart(2, "0");
  var minutes = date.getMinutes().toString().padStart(2, "0");
  var secounds = date.getSeconds().toString().padStart(2, "0");

  return year + "/" + month + "/" + day + " " + hours + ":" + minutes + ":" + secounds;
}

// ステータスを変換する関数
let formatStatus = (status) => {
  var span = document.createElement("span");
  span.classList.add("status-icon");
  if (status === "STARTED"){
    span.classList.add("started");
    span.textContent = "実行中";
  } else if (status === "FINISHED"){
    span.classList.add("finished");
    span.textContent = "完了";
  } else {
    span.textContent = "不明";
  }
  return span;
}

// 全件数を表示する
let displayTotalRecord = (totalNum) => {
  var totalRecords = document.getElementById("total_records");
  totalRecords.textContent = "全" + totalNum + "件"
}

// JSONデータからテーブルレコードを作成する関数
let displayRedcords = (records) => {
  var tbody = document.getElementById("table_body");
  // 既存データを削除
  while(tbody.firstChild){
    tbody.removeChild(tbody.firstChild)
  }
  // 新規データを表示
  records.forEach(record => {
    var tr = document.createElement("tr");
    // リポジトリID
    var td1 = document.createElement("td");
    td1.textContent = record.repository_id;
    tr.appendChild(td1);
    // リポジトリ名
    var td2 = document.createElement("td");
    td2.textContent = record.repository_name;
    tr.appendChild(td2);
    // RUN　ID
    var td3 = document.createElement("td");
    td3.textContent = record.run_id;
    tr.appendChild(td3);
    // 実行ステータス
    var td4 = document.createElement("td");
    td4.appendChild(formatStatus(record.status));
    tr.appendChild(td4);
    // 開始日時
    var td5 = document.createElement("td");
    td5.textContent = formatDate(record.started_at);
    tr.appendChild(td5);
    // 終了日時
    var td6 = document.createElement("td");
    td6.textContent = formatDate(record.finished_at);
    tr.appendChild(td6);

    tbody.appendChild(tr);
  });
}

// ページネーターを作成する関数
let displayPageNation = (totalNum) => {
  var navList = document.getElementById("page_list");
  // 既存データを削除
  while(navList.firstChild){
    navList.removeChild(navList.firstChild)
  }

  totalPageNum = Math.ceil(totalNum / pageSize);
  for (let page = 1; page <= totalPageNum; page++) {
    var li = document.createElement("li");
    li.classList.add("page-item");
    if(page == currentPageNum ){
      li.classList.add("active");
    }
    var a = document.createElement("a");
    a.href = "#";
    a.id = "page_" + page;
    a.classList.add("page-link");
    a.textContent = page;
    li.appendChild(a)
    navList.appendChild(li);

    // クリック時に選択したページで再表示させる
    a.addEventListener('click', (event) => {
      currentPageNum = document.getElementById(event.target.id).textContent;
      search();
    })
  }
}

// 表示する関数
let display = (data) => {
  displayRedcords(data.histories);
  displayPageNation(data.count);
  displayTotalRecord(data.count);
}

// 検索する関数
let search = () => {
  console.log("検索します")
  uri = createRequestUri();
  fetch(uri)
    .then((response) => response.json())
    .then((data) => display(data));
}

// 初期化関数
let initDashboard = () => {
  // イベントリスナー登録
  // 検索ボタン押下時に検索実行
  searchButton = document.getElementById("search_button");
  searchButton.addEventListener('click', search);
  // 実行履歴更新時に検索実行
  document.addEventListener('history_update', search);
  // 初期化処理
  search();
}

// イベントに対応する処理の追加
window.addEventListener('load', initDashboard);

