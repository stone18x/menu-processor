# General Setup 

Check https://firebase.google.com/docs/emulator-suite/install_and_configure and follow the setup instructions. 

# menu-processor Setup

First run the following command to log in to the Firebase CLI:

```
firebase login
```

Next run the following command to create a project alias. First, select `menu-processor` from the list. When asked what alias you want to use, choose default.

```
firebase use --add
```

The output should look like the following example. Remember to choose your actual Firebase project from the list:

```
? Which project do you want to add? YOUR_PROJECT_ID
? What alias do you want to use for this project? (e.g. staging) default

Created alias default for YOUR_PROJECT_ID.
Now using alias default (YOUR_PROJECT_ID)
```

To start the emulator and import some test data invoke: 

```
firebase emulators:start --import=./export
```