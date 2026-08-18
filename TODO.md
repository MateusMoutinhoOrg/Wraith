## Objective:
Create the Wraith financial brain.

## READ

- Todo-Specs/BrainSample/* 
  - description:The Whole Visualization of wraith in the user workdir.

- Todo-Specs/BrainSample/Help/*
  - description:Contains the guides of all the mechanics the project needs to have



### Expected Sandbox Tree:
- sandbox/cli 
  - action: Refatore aplying the new commands, and calling the new api 

- sandbox/config/
  - action: Refatore with these new configuration

- sandbox/lib/
  - action: Refatore aplying the new and the new contracts

- sandbox/contracts/deps/serverdeps
  - action: rename these module to requestdeps

- sandbox/Tasks/Tasks/
  - action: implement each task, one per file 

- sandbox/Tasks/run.go
  - action: implement the switcher to call eatch tasks based on TaskArray

- sandbox/Visualization/Visualization/
  - action: implement each visualization, one per file 

- sandbox/Visualization/run.go
  - implement the switcher to call eatch visualization based on VisualizationArray

