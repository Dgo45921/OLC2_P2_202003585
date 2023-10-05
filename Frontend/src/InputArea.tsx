import React, { useState, useRef } from "react";
import Editor from "@monaco-editor/react";
import { Tab, Tabs, TabList, TabPanel } from 'react-tabs';
import 'react-tabs/style/react-tabs.css';
import './InputArea.css';
import { Button } from 'react-bootstrap';
import ReactD3Tree from 'react-d3-tree';

function InputArea() {
    const [outputText, setOutputText] = useState("");

    const updateOutputText = (text:string) => {
        setOutputText(text);
    };
    const [tabs, setTabs] = useState([
        { language: "swift", code: "" },
    ]);
    const [currentTabIndex, setCurrentTabIndex] = useState(0);
    const handleTabAdd = () => {
        setTabs([...tabs, { language: "swift", code: "" }]);
    }

    const handleTabRemove = (index: number) => {
        const newTabs = tabs.filter((tab, i) => i !== index);
        setTabs(newTabs);
        setCurrentTabIndex(0);
    }

    const handleCodeChange = (value: string | undefined, index: number) => {
        const newTabs = [...tabs];
        if (typeof value === "string") {
            newTabs[index].code = value;
        }
        setTabs(newTabs);
    }

    const handleGuardar = () => {
        const currentTab = tabs[currentTabIndex];
        const fileContent = currentTab.code;
        const blob = new Blob([fileContent], { type: 'text/plain' });
        const url = URL.createObjectURL(blob);
        const link = document.createElement('a');
        link.href = url;
        link.download = 'codigo.swift';
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
    };

    const handleLogCode = async () => {



        
        const currentTab = tabs[currentTabIndex];
        console.log(currentTab.code);
        updateOutputText('')
        


    let data = {
            'code' : currentTab.code
        }


        try {
            const response = await fetch("http://localhost:5000/interpreter/parse", {
              method: "POST", // or 'PUT'
              headers: {
                "Content-Type": "application/json",
              },
              body: JSON.stringify(data),
            });
        
            const result = await response.json();

            
            
            updateOutputText(result.console)
          } catch (error) {
            console.error("Error:", error);
          }
        


        


    }

    const handleNewFile = () => {
        const newTabs = [...tabs];
        newTabs[currentTabIndex].code = "";
        setTabs(newTabs);
    };

    const handleTabSelect = (index: number) => {
        setCurrentTabIndex(index);
    }



    return (
        <div>
            <Button onClick={handleTabAdd}>Agregar pestaña</Button>
            <Tabs className="tabs" onSelect={handleTabSelect} selectedIndex={currentTabIndex}>
                <TabList className="tab-list">
                    {tabs.map((tab, index) => (
                        <Tab key={index} className="tab">
                            {"swift("}{index + 1}{")"}
                            <button
                                onClick={() => handleTabRemove(index)}
                                className="remove-tab-btn"
                            >
                                x
                            </button>
                        </Tab>
                    ))}
                </TabList>
                {tabs.map((tab, index) => (
                    <TabPanel key={index} className="tab-panel">
                        <Editor
                            height={"800px"}
                            width={"100%"}
                            theme={"vs-dark"}
                            language={tab.language}
                            value={tab.code || ''}
                            options={{
                                fontSize: 25, // Adjust the font size as needed
                                // Other editor options here
                            }}
                            onChange={(value) => handleCodeChange(value, index)}
                        />

                    </TabPanel>
                ))}
            </Tabs>

            <div className="button-container">
                <Button variant="success" onClick={handleNewFile} style={{ marginRight: 10 }}>Nuevo</Button>
                <label htmlFor="file-input" className="btn btn-warning" style={{ marginRight: 10 }}>Abrir</label>
                <input
                    id="file-input"
                    type="file"
                    accept=".swift"
                    style={{ display: "none" }}
                    onChange={(event) => {
                        const file = event.target.files?.[0];
                        if (file) {
                            const reader = new FileReader();
                            reader.onload = () => {
                                const fileContent = reader.result as string;
                                const newTabs = [...tabs];
                                newTabs[currentTabIndex].code = fileContent;
                                setTabs(newTabs);
                            };
                            reader.readAsText(file);
                        }
                    }}
                />
                <Button onClick={handleGuardar}>Guardar</Button>

                <Button onClick={handleLogCode} className="analyze-btn">Ejecutar</Button>
            </div>

            <h4 style={{color:"white", marginTop:"50px"}}>Salida:</h4>

            <textarea className="output-textarea" readOnly value={outputText} />


        </div>
    )
}

export default InputArea;
