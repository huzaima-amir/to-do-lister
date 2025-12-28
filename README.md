# to do lister

#### Application to manage tasks and events for users
Users can sign up using name, username and password, and login using username and password. Name, username and password can be changed. 
Users' tasks have deadlines, titles, and descriptions and can be "started" and "end" whenever the user marks them as such, and doing so updates the progress(pending, in progress or finished) accordingly. The overdue condition should switch to true if deadline passes and task isnt finished yet (handlers and routes for this function incomplete).
Users' events have start time and end time and the progress (upcomming, in progress or finished) updates according to time(handler and routes for this specific function incomplete). Events also have a location and can be marked as online or not online. 
Tags can be applied to both events and tasks to implement filtering. (filtering not implemented yet). Tags have a title and description.
All events and tasks can have a subtask checklist. These subtasks only have a title and can be checked or unchecked. Finishing a task should mass complete all subtasks. 

