# Todo_app
Todo_app: Created basic CRUD operations on local file system using NodeJs

API-1: Gets all the tasks from data.json file and the length of the data is given on the top

![Screenshot 2023-04-20 at 6 36 21 PM](https://user-images.githubusercontent.com/125876024/233376075-0cb27137-6933-4386-9d4a-0ea0837f8372.png)

API-2: Gets the specific task. By putting the task number at end point of API.

![Screenshot 2023-04-20 at 6 36 33 PM](https://user-images.githubusercontent.com/125876024/233377047-b3b86e6b-d4b5-4a7e-b103-1b8884be6774.png)


API-3: Creates new tasks. The id is auto-increamented each time when the user adds the new tasks.

![Screenshot 2023-04-20 at 6 36 48 PM](https://user-images.githubusercontent.com/125876024/233379563-fecf9eca-7f11-43bb-b5a5-d80ae131449f.png)


API-4: Updates the status of the task. here the validation is also added so as if the user gives 
the id which is not in the tasks list the error shown.

Successfull Updation of the task.

![Screenshot 2023-04-20 at 6 37 10 PM](https://user-images.githubusercontent.com/125876024/233380147-3356aaa9-c387-4eb8-94d3-097edda653ed.png)


Error if the task id not in the list.

![Screenshot 2023-04-20 at 6 55 40 PM](https://user-images.githubusercontent.com/125876024/233380542-fd69762c-8fa4-4a5c-bce5-e197ec245dec.png)


API-5: Deleting the tasks by id specified in the end point of the API. Here the confirmation of the deleted task is given by the 204.

![Screenshot 2023-04-20 at 6 37 55 PM](https://user-images.githubusercontent.com/125876024/233381390-4ca0428b-26cc-48e9-ae43-2b660f2b5519.png)



Versions:
"Express": "^4.18.2" [latest ones can give fatal error- hence use major version as 4]
"nodemon": "^2.0.22"


