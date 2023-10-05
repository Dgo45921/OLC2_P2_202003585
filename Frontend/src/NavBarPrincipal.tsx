import 'bootstrap/dist/css/bootstrap.css';
import * as ReactBootstrap from 'react-bootstrap';
import * as d3 from 'd3-graphviz'



async function errorReport() {
    const response = await fetch("http://localhost:5000/interpreter/getErrors");
    const textData = await response.text();
    const jsonbody = JSON.parse(textData);
  
    // Create a new div element to serve as the container for the graph
    const graphContainer = document.createElement("div");
    console.log(jsonbody.dotCode)
  
    // Render the graph using d3-graphviz
    d3.graphviz(graphContainer)
      .renderDot(jsonbody.dotCode)
      .on("end", () => {
        // Create a new HTML document for the new tab
        const newTab = window.open("", "_blank");
        if (newTab) {
          // Append the graph container to the new document's body
          newTab.document.body.appendChild(graphContainer.cloneNode(true));
        } else {
          console.error("Failed to open new tab");
        }
      });

      
  }

  async function ASTreport() {
    await fetch("http://localhost:5000/interpreter/getCST");
  }
  
  
  
  


async function SymbolTableReport() {
  const response = await fetch("http://localhost:5000/interpreter/getST");
  const textData = await response.text();
  const jsonbody = JSON.parse(textData);

  // Create a new div element to serve as the container for the graph
  const graphContainer = document.createElement("div");
  console.log(jsonbody.dotCode)

  // Render the graph using d3-graphviz
  d3.graphviz(graphContainer)
    .renderDot(jsonbody.dotCode)
    .on("end", () => {
      // Create a new HTML document for the new tab
      const newTab = window.open("", "_blank");
      if (newTab) {
        // Append the graph container to the new document's body
        newTab.document.body.appendChild(graphContainer.cloneNode(true));
      } else {
        console.error("Failed to open new tab");
      }
    });
}


function NavBarPrincipal() {


    return (
        <ReactBootstrap.Navbar bg="dark" variant="dark" style={{marginBottom: 10}}>
            <ReactBootstrap.Navbar.Brand style={{marginLeft: 50}}>T - Swift</ReactBootstrap.Navbar.Brand>
            <ReactBootstrap.Nav className="mr-auto">
                <ReactBootstrap.NavDropdown title="Reportes" id="reportes-dropdown">
                    <ReactBootstrap.NavDropdown.Item onClick={errorReport}>Reporte de errores</ReactBootstrap.NavDropdown.Item>
                    <ReactBootstrap.NavDropdown.Item onClick={ASTreport}>Reporte CST</ReactBootstrap.NavDropdown.Item>
                    <ReactBootstrap.NavDropdown.Item onClick={SymbolTableReport}>Reporte Tabla de símbolos</ReactBootstrap.NavDropdown.Item>
                </ReactBootstrap.NavDropdown>
            </ReactBootstrap.Nav>
        </ReactBootstrap.Navbar>
    );
}


export default NavBarPrincipal
