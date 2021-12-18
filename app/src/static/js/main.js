function toggle_view(id_name) {
    var x = document.getElementById(id_name);
    if (x.style.display === "none") {
      x.style.display = "block";
    } else {
      x.style.display = "none";
    }
  }
  
document.getElementById("monthly_button").onclick = function() {toggle_view("monthly_div")};
  
document.getElementById("long_term_button").onclick = function() {toggle_view("long_term_div")};
  
document.getElementById("planning_button").onclick = function() {toggle_view("planning_div")};