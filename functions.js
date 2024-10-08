
function reset() {
    var input = document.getElementById("mtg-search");
    input.value = "";
    filterCards();
}

function filterCards() {
    var input = document.getElementById("mtg-search");
    var filter = input.value.toUpperCase();
    var table = document.getElementById("mtg-table");
    var tr = table.getElementsByTagName("tr");

    table.style.display = "none";
    // Starting at 1 to ignore the headers
    for (var i = 1; i < tr.length; i++) {
        var td = tr[i].getElementsByTagName("td");
        if (td) {
            var foundMatch = false;
            for (var j = 0; j < td.length; j++) {
                var txtValue = td[j].textContent || td[j].innerText;
                if (txtValue.toUpperCase().indexOf(filter) > -1) {
                    foundMatch = true;
                }
            }
            if (foundMatch) {
                tr[i].style.display = "";
            } else {
                tr[i].style.display = "none";
            }
        }
    }
    table.style.display = "";
}

