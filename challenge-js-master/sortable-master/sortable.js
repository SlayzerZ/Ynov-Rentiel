async function LoadInTable(url,table){
    const power = ["Combat : ", "Durability : ", "Intelligence : ", "Power : ", "Speed : ", "Strength : "]
    const tableau = ["Icon","Name","Full Name","Powerstats","Race","Gender","Height","Weight","Place Of Birth","Alignement"]
    const tr = "<tr></tr>"
    const Thead = document.querySelector("thead")
    const Tbody = document.querySelector("tbody")
    const reponse = await fetch(url)
    const data = await reponse.json()
    Thead.innerHTML = tr + tr
    Tbody.innerHTML = tr
    let n = 0
    let b = 0
    
    const title = document.createElement("th")
    title.colSpan="10"
    title.textContent = "Comics"
    title.className = "title"
    Thead.querySelectorAll("tr")[0].appendChild(title)
    let select = document.getElementById("Element-select")
    select.addEventListener('change', function () {
        select = this.value;
        console.log(select)
    })
    document.querySelector('#nextButton').addEventListener('click', function () {
        if((curPage * select) < data.length) curPage++;
        renderTable();
    })
    document.querySelector('#prevButton').addEventListener('click', function () {
        if(curPage > 1) curPage--;
        renderTable();
    })
    while(n < tableau.length) {
        const stitle = document.createElement("th")
        stitle.textContent = tableau[n]
        Thead.querySelectorAll("tr")[1].append(stitle)
        n += 1
    }
    n = 0
    while(n < data.length) {
        b = 0
        const imagef = document.createElement("th")
        imagef.className = "icon"
        Tbody.querySelectorAll("tr")[n].append(imagef)
        const imaget = document.createElement("img")
        imaget.src = data[n].images.xs
        Tbody.querySelectorAll("th.icon")[n].appendChild(imaget)

        const name = document.createElement("th")
        name.textContent = data[n].name
        Tbody.querySelectorAll("tr")[n].append(name)

        const Fname = document.createElement("th")
        Fname.textContent = data[n].biography.fullName
        Tbody.querySelectorAll("tr")[n].append(Fname)
        
        const Powerstats = document.createElement("th")
        Powerstats.className = "stats"
        Tbody.querySelectorAll("tr")[n].append(Powerstats)
        while (b < power.length) {
            const Powerstats = document.createElement("p")
            if (power[b] === power[0]) {
            Powerstats.textContent = power[b] + data[n].powerstats.combat
            }
            if (power[b] === power[1]) {
                Powerstats.textContent = power[b] + data[n].powerstats.durability
            }
            if (power[b] === power[2]) {
                Powerstats.textContent = power[b] + data[n].powerstats.intelligence
            }
            if (power[b] === power[3]) {
                Powerstats.textContent = power[b] + data[n].powerstats.power
            }
            if (power[b] === power[4]) {
                Powerstats.textContent = power[b] + data[n].powerstats.speed
            }
            if (power[b] === power[5]) {
                Powerstats.textContent = power[b] + data[n].powerstats.strength
            }
            Tbody.querySelectorAll("th.stats")[n].append(Powerstats)
            b += 1
        }


        const race = document.createElement("th")
        race.textContent = data[n].appearance.race
        Tbody.querySelectorAll("tr")[n].append(race)

        const gender = document.createElement("th")
        gender.textContent = data[n].appearance.gender
        Tbody.querySelectorAll("tr")[n].append(gender)

        const height = document.createElement("th")
        height.textContent = data[n].appearance.height
        Tbody.querySelectorAll("tr")[n].append(height)

        const weight = document.createElement("th")
        weight.textContent = data[n].appearance.weight
        Tbody.querySelectorAll("tr")[n].append(weight)

        const pob = document.createElement("th")
        pob.textContent = data[n].biography.placeOfBirth
        Tbody.querySelectorAll("tr")[n].append(pob)

        const alignment = document.createElement("th")
        alignment.textContent = data[n].biography.alignment
        Tbody.querySelectorAll("tr")[n].append(alignment)

        Tbody.innerHTML += tr
        n += 1
    }
    console.log(select)
}
  function renderTable() {
    // create html
    let result = '';
    data.filter((row, index) => {
          let start = (curPage-1)*pageSize;
          let end =curPage*pageSize;
          if(index >= start && index < end) return true;
    }).forEach(c => {
       result += `<tr>
       <td>${c.name}</td>
       <td>${c.age}</td>
       <td>${c.breed}</td>
       <td>${c.gender}</td>
       </tr>`;
    });
    table.innerHTML = result;
  }
LoadInTable('https://rawcdn.githack.com/akabab/superhero-api/0.2.0/api/all.json', document.querySelector("table"))