# VereinsfinanzManager - VfM
Ein Go-Backend Projekt im Rahmen der Vorlesung Backend-Entwicklung im Studiengang Professionale Softwarae Engineering (PSE) zum Master an der Knowledge Foundation @ Reutlingen University
## The current json of our Club are as following

### Club
````json
{
  "schema_version": "1",
  "club_name": "TSV Germania Hackentrick",
  "club_leader": "Karl-Heinz Rummenigge",
  "budget": 20000000,
  "address": "Hardtstraße 37, 76185 Karlsruhe",
  "description": "Der regionale Club mit vielen Möglichkeiten und einer tollen Atmosphäre. Insights you can play on.",
  "bank_account": 
    {
      "schema_version": "1",
      "bank_account_id": 1,
      "owner_name": "Karl-Heinz Rummenigge",
      "name_of_bank": "Bank xyz",
      "iban": "DE123456765432345678"
    },
  "departments_id": []
}
````

## The current json of our Departments are as following

### Handball
```json
{
  "schema_version": "1",
  "name_of_department": "Handball",
  "department_leader": "Heiner Brandt",
  "department_budget": 5000000,
  "department_cost":4880000,
  "teams_id": []
}
```
### Fußball
````json
{
  "schema_version": "1",
  "name_of_department": "Fußball",
  "department_leader": "Hansi Flick",
  "department_budget": 10000000,
  "department_cost":8800000,
  "teams_id": []
}
````
### Schwimmen
````json
{
  "schema_version": "1",
  "name_of_department": "Schwimmen",
  "department_leader": "Franziska van Almsick",
  "department_budget": 100000,
  "department_cost":100000,
  "teams_id": []
}
````

## The current json of our Teams are as following

### Handball Teams
```json
{
  "schema_version": "1",
  "name_of_team": "Rhein-Neckar Löwen 1",
  "team_leader": "Markus Strobel",
  "team_budget": 2550000,
  "overall_costs": 1000
}
```
```json
{
  "schema_version": "1",
  "name_of_team": "Rhein-Neckar Löwen 2",
  "team_leader": "Martin Schwalb",
  "team_budget": 1550000,
  "overall_costs": 1000
}
```
```json
{
  "schema_version": "1",
  "name_of_team": "Rhein-Neckar Löwen 3",
  "team_leader": "Roland Mächtel",
  "team_budget": 500000,
  "overall_costs": 1000
}
```
### Fussball Teams
```json
{
  "schema_version": "1",
  "name_of_team": "Rhein-Neckar Kickers 1",
  "team_leader": "Joachim Löw",
  "team_budget": 6500000,
  "overall_costs": 1000
}
```
```json
{
  "schema_version": "1",
  "name_of_team": "Rhein-Neckar Kickers 2",
  "team_leader": "Thomas Tuchel",
  "team_budget": 2500000,
  "overall_costs": 1000
}
```
```json
{
  "schema_version": "1",
  "name_of_team": "Rhein-Neckar Kickers 3",
  "team_leader": "Pep Guardiola",
  "team_budget": 1000000,
  "overall_costs": 1000
}
```
## The current json of our non / human resources are as following 
### Human Resources
````json
{
  "schema_version": "1",
  "name": "Kai Havertz",
  "value": 9000000,
  "salary": 5000000,
  "contract_runtime": "2 years",
  "team_id": ""
}
````

### Non Human Resources 
````json
{
  "schema_version": "1",
  "name": "Lizenzgebühr 1. Liga",
  "cost": 1500000,
  "validity": "Gültig bis 2022",
  "duration": "2 years",
  "team_id": ""
}
````

The Presentation of the Project is available on Google Drive: 
###### https://docs.google.com/presentation/d/1xUGWtlGYvSkyTbJNTZdRMbv3pCi4z6LUkTdBgEv8BuQ/edit?usp=sharing 
The final presentation will be uploaded in the repo as a powerpoint .pptx
A$AP

- by Marvin Bermel & Mikka Jenne
