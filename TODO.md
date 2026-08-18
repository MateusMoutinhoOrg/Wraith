## Objective:
Create the Wraith financial brain.

## READ

- WraithBrainSample/* 
  - description:The Whole Visualization of wraith in the user workdir.

- WraithBrainSample/Help/*
  - description:Contains the guides of all the mechanics the project needs to have


### IMMPORTANT:
- task and visualization must be easly changed, since the are the. heart of the project
- the doc must be easier to beguiner,since the ideia of these project, its to allow people create or use a "second brain"

### Expected Sandbox Tree:
- sandbox/cli 
  - action: Refatore aplying the new commands, and calling the new api 
- sandbox/config/
  - action: Refatore with these new configuration

- sandbox/lib/
  - action: Refatore aplying the new and the new contracts
- sandbox/contracts/api/api.go
  - model: use sandbox/contracts/api/api.go as the model of how it needs to be made the contract 

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


### Expected Doc Organization: 

### Brain-Usage
- these doc its the actual Cli-Usage (needs to rename), and must cover every thing about how to execute actions, show visualizations,etc..

### Brain-Config
- explain how to add new tasks, and new visualizations in the given brain,  it needs to show how to fork these template, rename to the new github name, and configure the users functions.

## Developer
- just refatore to these project, but the ideia. is the same

## Templating
 - will be removed, all its contents will go to Brain-Config ,since the ideia of these projetct, its to be forked.