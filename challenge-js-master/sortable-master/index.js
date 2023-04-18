document.addEventListener('DOMContentLoaded', init);

let select = document.getElementById("Element-select")
    select.addEventListener('change',function () {
        select.value = this.value;
        // console.log(select.value)
        pageSize = select.value
        createTable()
    })
let data, table, sortCol;
let sortAsc = false; 
let pageSize = select.value; // Nombre de personnages par page
let curPage = 1; // Index de la page courante
let autre;

async function init() {
  table = document.querySelector('#ComicsTable tbody'); // Sélection de notre tableau
  let resp = await fetch('https://rawcdn.githack.com/akabab/superhero-api/0.2.0/api/all.json');
  data = await resp.json();
  for (let i = 0; i < data.length; i ++){
    data[i].fullname = data[i].biography.fullName
    data[i].race = data[i].appearance.race
    data[i].gender = data[i].appearance.gender
    data[i].height = data[i].appearance.height[1]
    data[i].weight = data[i].appearance.weight[1]
    data[i].placeofbirth = data[i].biography.placeOfBirth
    data[i].alignment = data[i].biography.alignment
  }
  
  createTable();
  

  // Ajout des écouteurs d'événements pour le tri et la pagination
  document.querySelectorAll('#ComicsTable thead tr th').forEach(t => {
    t.addEventListener('click', sort);
  });

  document.querySelector('#prev').addEventListener('click', previousPage);
  document.querySelector('#next').addEventListener('click', nextPage);
}

function createTable() {
  let b = -1
  let result = '';
  data.filter((row, index) => {
    let start = (curPage - 1) * pageSize;
    let end = curPage * pageSize;
    if (index >= start && index < end) return true;
  }).forEach(i => {
    b++
    result += `<tr>
     <td><img src=${i.images.xs}></td>
     <td>${i.name}</td>
     <td>${i.biography.fullName}</td>
     <td>Combat: ${i.powerstats.combat}<br>
     Durabilité: ${i.powerstats.durability}<br>
     Intelligence: ${i.powerstats.intelligence}<br>
     Puissance: ${i.powerstats.power}<br>
     Vitesse: ${i.powerstats.speed}<br>
     Force: ${i.powerstats.strength}<br>
     </td>
     <td>${i.appearance.race}</td>
     <td>${i.appearance.gender}</td>
     <td>${i.appearance.height[1]}</td>
     <td>${i.appearance.weight[1]}</td>
     <td>${i.biography.placeOfBirth}</td>
     <td>${i.biography.alignment}</td>
     </tr>`;
  });
  table.innerHTML = result;
}

function sort(e) {
  const thisSort = e.target.dataset.sort;
  let sortAscNew = true;

  if (sortCol === thisSort) {
    sortAscNew = !sortAsc;
  } else {
    sortCol = thisSort;
  }

  data.sort((a, b) => {
    let sortResult = 0;
    
    if (a[sortCol]< b[sortCol]) {
      sortResult = sortAscNew ? 1 : -1;
    } else if (a[sortCol] > b[sortCol]) {
      sortResult = sortAscNew ? -1 : 1;
    }
    if (a[sortCol] === "") {
      sortResult = sortAscNew ? 1 : 1;
    }
    if (a[sortCol] === "-") {
      sortResult = sortAscNew ? 1 : 1;
    }

    return sortResult;
    
  });

  sortAsc = sortAscNew;
  createTable();
}

function previousPage() {
  if (curPage > 1) curPage--;
  createTable();
}

function nextPage() {
  if ((curPage * pageSize) < data.length) curPage++;
  createTable();
}

