
### Brain Workflow:

## Process:

### Show Error 
- creates a Error.md containing informations about the error.

## Glossary:

## Tick Worflow:
- User runs: 
~~~bash
./brain tick
~~~
- For each project:
  - brain trys to read: <Project>/Task.yaml
  - if the file does not exist:
      - shows a error 
      - stop execution.
  - if   <Project>/Task.yaml.name not contains a valida action:
      - shows a error 
      - stop execution.  
  - executes  <Project>/Task.yaml.name action.
  - if <Project>/Task.yaml.name action fails:
    - shows a error
    - stop execution.

  - renders all markdowns of <Project>/Dasboard


