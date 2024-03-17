# Functional Specifications - Coding Dojo

## Document Control
**Document Information:**
| | Information |
| --- | --- |
| Document Owner | Maxime CARON |
| Creation Date | 2024-03-17 |
| Last Updated | 2024-03-17 |
| Document Name | Functional Specifications - Coding Dojo |

**Version History:**
| Version | Date | Author | Description |
| --- | --- | --- | --- |
| 0.1 | 2024-03-17 | Maxime CARON | Initial draft |


## Summary
<details>

<summary>Click to expand!</summary>

- [Functional Specifications - Coding Dojo](#functional-specifications---coding-dojo)
  - [Document Control](#document-control)
  - [Summary](#summary)
  - [Overview](#overview)
    - [Scope](#scope)
    - [Out of Scope](#out-of-scope)
  - [Customers](#customers)
    - [For version 1.0](#for-version-10)
    - [For future versions](#for-future-versions)
  - [Use Cases](#use-cases)
    - [Case 1 - Student wants to learn algorithmic concepts](#case-1---student-wants-to-learn-algorithmic-concepts)
    - [Case 2 - Teacher wants to target student problems](#case-2---teacher-wants-to-target-student-problems)
    - [Case 3 - Student Council wants to organize events and competitions](#case-3---student-council-wants-to-organize-events-and-competitions)
    - [Case 4 - Student want to propose new exercises](#case-4---student-want-to-propose-new-exercises)
    - [Case 5 - Studeant want to evaluate his skills](#case-5---studeant-want-to-evaluate-his-skills)
  - [Constraints](#constraints)
  - [UX / UI](#ux--ui)
    - [User Experience](#user-experience)
    - [User Interface](#user-interface)
  - [Main features](#main-features)
    - [Exercises](#exercises)
    - [Documentation](#documentation)
    - [Reward system](#reward-system)
  - [Security](#security)
  - [Success Criteria's](#success-criterias)
  - [Non-Functional Requirements](#non-functional-requirements)
  - [Future Improvements](#future-improvements)
  - [Glossary](#glossary)

</details>

## Overview
Every day, new people try to discover the world of programming, and every day, some of them are confronted with a similar problem. They can learn the basics, but they have trouble learning the algorithms. That's why this project exists. This project is a web application designed to help new developers acquire increased programming skills. With exercises and lessons provided, beginners will be capable of entering the world of algorithms step by step.

### Scope
This project aims to help people progress in algorithmic concepts and acquire new programming skills, but also to help them evaluate their current skills.

### Out of Scope
The project is intended as a learning platform, but it is clearly not intended to replace schools and teachers. It will not deliver any kind of diploma or certification.

## Customers
### For version 1.0

| Stakeholder | Description |
| --- | --- |
| ALGOSUP students | For the first version, the principal customer will be ALGOSUP students, who will exploit the system to improve their skills and gradually learn algorithmic concepts. They can also suggest new exercises on new concepts if they think something is missing. |
| ALGOSUP teachers | Teachers will be able to use the app to target student problems more precisely and focus on them. |
| ALGOSUP Student Council | The student council will be able to use the app to organize events and competitions. |

### For future versions

| Stakeholder | Description |
| --- | --- |
| New developers | The app will be open to everyone, and new developers will be able to use it to acquire new programming skills. |
| Companies | Companies will be able to use the app to evaluate the skills of their future employees. |

## Use Cases
### Case 1 - Student wants to learn algorithmic concepts
| Name | Brief description | Behavior |
| --- | --- | --- |
| Lucas | Lucas is in his first year at ALGOSUP and wants to learn more about algorithmic concepts. | Lucas logs in to the app and starts with the first exercise. |

| Step | Description | Expected Result |
| --- | --- | --- |
| 1 | Lucas logs in to the app. | Lucas is redirected to the home page. |
| 2 | Lucas selects the first exercise. | Lucas is redirected to the exercise page. |
| 3 | Lucas reads rules and given documentation. | Lucas understands the exercise. |
| 4 | Lucas do the exercise. | Lucas folows the instructions and finishes the exercise. |
| 5 | Lucas gets a reward. | Lucas is happy and motivated to continue. |

### Case 2 - Teacher wants to target student problems
| Name | Brief description | Behavior |
| --- | --- | --- |
| Mr. Smith | Mr. Smith is a teacher at ALGOSUP and wants to target student problems more precisely. | Mr. Smith logs in to the app and checks the students' progress. |

| Step | Description | Expected Result |
| --- | --- | --- |
| 1 | Mr. Smith logs in to the app. | Mr. Smith is redirected to the home page. |
| 2 | Mr. Smith selects the students' progress page. | Mr. Smith is redirected to the students' progress page. |
| 3 | Mr. Smith checks the students' progress. | Mr. Smith can see the students' progress. |
| 4 | Mr. Smith targets student problems. | Mr. Smith can see the students' problems. |
| 5 | Mr. Smith focuses on them. | Mr. Smith can focus on the students' problems. |

### Case 3 - Student Council wants to organize events and competitions
| Name | Brief description | Behavior |
| --- | --- | --- |
| Student Council | The student council wants to organize events and competitions. | The student council logs in to the app and organizes events and competitions. |

| Step | Description | Expected Result |
| --- | --- | --- |
| 1 | The student council logs in to the app. | The student council is redirected to the home page. |
| 2 | The student council selects the events and competitions page. | The student council is redirected to the events and competitions page. |
| 3 | The student council organizes events and competitions. | The student council can organize events and competitions. |
| 4 | The student council invites students. | The student council can invite students. |
| 5 | The student council can see the students' progress. | The student council can see the students' progress. |

### Case 4 - Student want to propose new exercises
| Name | Brief description | Behavior |
| --- | --- | --- |
| Julia | Julia knows a new concept and wants to propose a new exercise. | Julia logs in to the app and proposes a new exercise. |

| Step | Description | Expected Result |
| --- | --- | --- |
| 1 | Julia logs in to the app. | Julia is redirected to the home page. |
| 2 | Julia selects the propose new exercise page. | Julia is redirected to the propose new exercise page. |
| 3 | Julia proposes a new exercise. | Julia describes the new exercise. |
| 4 | Julia submits the new exercise. | Julia submits the new exercise. |
| 5 | Julia gets a reward. | Julia is happy and motivated to continue. |
| 6 | The exercise is analyzed. | The exercise is analyzed by the team. |
| 7 | The exercise is added. | The exercise is added to the app. |

### Case 5 - Studeant want to evaluate his skills
| Name | Brief description | Behavior |
| --- | --- | --- |
| John | John wants to evaluate his skills. | John logs in to the app and evaluates his skills. |

| Step | Description | Expected Result |
| --- | --- | --- |
| 1 | John logs in to the app. | John is redirected to the home page. |
| 2 | John selects his profile. | John is redirected to his profile page. |
| 3 | John can see his progress. | John can see his progress through statistics and diagrams. |

## Constraints
The main constraints are the following:
- Time:  the prototype must be delivered in 3 months, and the first version in 6 months.
- Budget: the budget is close to 0€.
- Privacy: the app must respect user privacy.
- Accessibility: the app must be accessible from any web browser.
- Maintenance: the app must be straightforward to maintain and update.
- Differentiation: the app must differentiate ALGOSUP students from other users.

## UX / UI

### User Experience
User experience is a critical part of the project. The app must be easy to use and understand. The user must be able to navigate through the app without any difficulty. The app must be accessible from any web browser.

### User Interface
**Graphic Charter:**
The style must respect the following graphic charter: [here](../graphic-charter/graphic-charter.md)

**Pages:**
The app will contain the following pages:
| Page | Description | Content |
| --- | --- | --- |
| Connection / Registration | The connection and registration page will contain the form to connect or register. | - Username<br>- Password<br>- Register button<br>- Connection button |
| Home | The home page will contain the main features of the app. | - Exercises<br>- Documentation<br>- Reward system |

## Main features
### Exercises

### Documentation

### Reward system

## Security

## Success Criteria's

## Non-Functional Requirements

## Future Improvements

## Glossary
