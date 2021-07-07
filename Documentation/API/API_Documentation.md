# Here is our documentation of our API functions and commands
The following commands are available: 
Firstly you can create a Club to start the basics of the application.
This runs with a simple Post-API Request with the club information as json.

	// clubs
	router.HandleFunc("/club", handler.CreateClub).Methods("POST")

To handle the club of your own you can get it with a special id. This you can handle with a simple GET-API Request: 

	router.HandleFunc("/clubs", handler.GetClubs).Methods("GET")
	router.HandleFunc("/clubs/{id}", handler.GetClub).Methods("GET")

To Update information of your club you can start an drun a PUT-API Command with all updated information which will request as json:

	router.HandleFunc("/clubs/{id}", handler.UpdateClub).Methods("PUT")

Finally, you can delete the club and all dependencies e.g. departments and teams of it:

	router.HandleFunc("/clubs/{id}", handler.DeleteClub).Methods("DELETE")

Partially same with the departments of your club. To add departments to xour club and to the application you can Add some with the POST-API command: 

    // departments
	router.HandleFunc("/department", handler.CreateDepartment).Methods("POST")

To request all departments, or a single department with an id you can run API-Calls as following: 

	router.HandleFunc("/departments", handler.GetDepartments).Methods("GET")
	router.HandleFunc("/departments/{id}", handler.GetDepartment).Methods("GET")

If you would like to change some information of your department or add some depending teams on the department you cann run a simple update with the PUT-API Call:
With this you create a json which will walk through and will compare to the existing one.

	router.HandleFunc("/departments/{id}", handler.UpdateDepartment).Methods("PUT")

At the end you are able to delete the department wit de DELETE API Call. With this command, you update the club and delete all teams who depend on the selected department.

	router.HandleFunc("/departments/{id}", handler.DeleteDepartment).Methods("DELETE")

At least you can create teams with all the information it contains. 

	// teams
	router.HandleFunc("/team", handler.CreateTeam).Methods("POST")

Like the other options and commands, you can find all teams in you organization with the simple GET function, or you look after a specific teams with a concrete ID:

	router.HandleFunc("/teams", handler.GetTeams).Methods("GET")
	router.HandleFunc("/teams/{id}", handler.GetTeam).Methods("GET")

If it gives some updates of the team you can update the information like the yearly costs of a team as well. 

	router.HandleFunc("/teams/{id}", handler.UpdateTeam).Methods("PUT")

To get the full access of all functions you can delete a team. With this you update also the department and the budget of the department as well. 

	router.HandleFunc("/teams/{id}", handler.DeleteTeam).Methods("DELETE")

<mark>To handle all financial questions an transactions we will continue with the splitting and grounding off all information insights a team e.g. contracts yearly, daily or weekly costs and things like that</mark>