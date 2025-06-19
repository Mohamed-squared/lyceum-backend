Here is the comprehensive user guide converted to Markdown format.

# Lyceum Learning Platform - Comprehensive User Guide

---

## Table of Contents

1.  **Introduction to Lyceum**
    *   What is Lyceum?
    *   Who is this Guide For?
    *   Key Features Overview
2.  **Getting Started**
    *   Accessing Lyceum
    *   Creating Your Account (Sign Up)
    *   Logging In
    *   Password Reset
    *   Initial Onboarding (First-time Setup)
3.  **Navigating Lyceum: The Main Interface**
    *   The Sidebar
        *   Standard Navigation
        *   Course-Specific Navigation
        *   Theme Toggle
    *   Main Content Area
    *   Header Bar & User Section
    *   Mobile Navigation
    *   Background & Opacity Customization (via Settings)
4.  **Home Dashboard**
    *   Overview
    *   Test Generation Snapshot
    *   Courses Snapshot
5.  **TestGen Suite: Personalized Test Generation**
    *   TestGen Dashboard
    *   Managing Subjects (for TestGen)
        *   Understanding Subject Details
        *   Selecting Your Active Subject
    *   Managing Studied Chapters
    *   Generating a New Test
        *   Testing Studied Chapters
        *   Testing Specific Chapters
        *   Configuring Your Test (Question Counts, Ratios, Timing)
        *   Starting an Online Test
        *   Generating a PDF Test (and Solutions)
        *   Generating a LaTeX Source (.tex)
    *   Taking an Online Test (TestGen)
        *   The Test Interface
        *   Answering MCQs
        *   Answering Problems (LaTeX Input)
        *   Navigating Questions
        *   Timer and Submission
    *   TestGen Exams Dashboard & History
        *   Viewing Pending PDF Exams
        *   Entering Results for PDF Exams
        *   Viewing Completed Exam History
        *   Reviewing Exam Details & AI Feedback (TestGen Exams)
    *   TestGen Progress Dashboard
        *   Understanding Charts (Difficulty, Mastery, Attempts, Wrongs)
6.  **Courses Platform**
    *   Browsing Courses
        *   Searching and Filtering
    *   Viewing Course Details
        *   Description, Tags, Chapters, Prerequisites, Corequisites
    *   Enrolling in a Course & Setting Your Pace
        *   Full Enrollment vs. Viewer Mode
    *   My Courses Dashboard
    *   Course-Specific Dashboard
        *   Today's Objective & Next Task
        *   Quick Links (Study Material, Assignments, Notes, Progress)
        *   Unenrolling from a Course
    *   Study Material
        *   Navigating Chapters and Videos
        *   YouTube Video Player
        *   Video Transcription & Interaction
        *   PDF Viewer & Interaction
        *   AI Tools for Study Material (Explain Selection, Ask AI)
        *   Chapter Formula Sheet (AI Generated)
        *   Chapter Summary (AI Generated)
        *   Skip Exam (Testing out of a Chapter)
    *   Assignments & Exams (Course-Specific)
        *   Accessing Assignments and Exams
        *   Taking Course Activities (Online Interface similar to TestGen)
        *   Viewing History for Course Activities
    *   Course Progress Details
        *   Overall Grade and Performance
        *   Performance Charts
        *   Score Breakdown
        *   Course Certificate (Generation & Download)
    *   Notes & Documents (Per Course, Per Chapter)
        *   Accessing the Notes Panel
        *   Adding Text & LaTeX Notes
        *   Uploading Files (PDF, Images, Text)
        *   Viewing and Editing Notes
        *   AI Tools for Notes (Review, Convert to LaTeX, Improve)
        *   Downloading Notes (PDF, TeX)
7.  **AI Chat Studio**
    *   Interface Overview (Sessions Sidebar, Chat Area)
    *   Starting a New Chat
    *   Sending Messages & Using LaTeX
    *   Managing Chat Sessions (Switching, Deleting)
8.  **Community & Engagement**
    *   Leaderboard
    *   Global Chat (Experimental)
        *   Accessing and Using Global Chat
        *   Sending Messages, Mentions (@username, @everyone)
        *   Replying to Messages
        *   Pinning/Unpinning Messages (Admin)
        *   Viewing Pinned Messages
        *   Deleting Messages (Owner/Admin)
    *   Marketplace (Coming Soon)
9.  **Your Profile & Settings**
    *   Viewing Your Profile
        *   Avatar, Display Name, Username, Email, Credits, Badges
    *   Editing Your Profile (Display Name, Avatar via Cropper)
    *   Settings Panel
        *   Profile & Account Management (Changing Username, Email, Password)
        *   Appearance Settings (Backgrounds, Card Opacity)
        *   Experimental Features (Toggling Chat, Marketplace, Music)
        *   AI Model & System Prompt Settings (Advanced Users)
    *   Logging Out
    *   Deleting Your Account
10. **Music & Sounds Player (Experimental)**
    *   Accessing the Music & Sounds Center
    *   Music Sources (YouTube Music)
        *   Playing YouTube URLs/Playlists
    *   Ambient Sounds
    *   Study Music (Lo-fi, Chill, etc.)
    *   Binaural Beats
    *   Main Player Controls
        *   Play/Pause, Next/Previous, Shuffle, Repeat
        *   Volume and Progress
    *   Current Playlist Management
        *   Saving and Loading User Playlists
    *   Mini Player (Floatie)
    *   Sound Settings (UI Sounds, Ambient Volume)
11. **Import / Export Test Data**
    *   Exporting Your TestGen Data
    *   Importing TestGen Data
12. **Admin Panel (For Administrators Only)**
    *   Accessing the Admin Panel
    *   Admin Navigation Sidebar
    *   Overview Section
        *   Platform Statistics
        *   Sending Global Announcements
    *   User Management
        *   Listing and Searching Users
        *   Viewing User Details
        *   Toggling Admin Status (Primary Admin)
        *   Changing Usernames (Primary Admin)
        *   Managing User Course Progress (Marking Complete, Setting Grade)
        *   Managing User Badges (Adding/Removing)
        *   Managing User-Submitted TestGen Subjects (Approving/Rejecting)
    *   Global Course Management (Administering /courses Collection)
        *   Viewing Courses Requiring Attention (Pending, Reported)
        *   Approving, Rejecting, Deleting Global Courses
        *   Adding New Global Courses
        *   Editing Existing Global Course Definitions
        *   Playlist & Chapter Resource Assignment
    *   Moderation (Feedback & Exam Issues)
        *   Viewing Recent Feedback and Exam Issues
        *   Replying to Users
        *   Deleting Feedback/Issues
        *   Deleting ALL Feedback & Issues
    *   Music & Sound Library Management (Admin)
        *   Managing Ambient, Study, Binaural Libraries
        *   Managing UI Sound Effects
        *   Generating Config Code for config.js
    *   Database Management (Advanced Admin)
        *   Browsing Collections and Documents
        *   Viewing and Editing Document Fields
        *   Deleting Documents
        *   Navigating Subcollections
    *   System Operations (Primary Admin)
        *   Admin Tasks / To-Do List
        *   Global AI System Prompts Management
        *   Chat Auto-Deletion Settings
        *   Course Exam Defaults Configuration
        *   Deleting User-Generated Content (Formula Sheets, Summaries)
    *   Admin Testing Aids
        *   Selecting User to Modify
        *   TestGen Subject Aids (Mark All Studied, Reset Progress)
        *   Course Aids (Mark Chapter Studied, Complete Activity, Set Status/Grade, Adjust Credits)
        *   Simulating Days Passed in a Course
13. **Troubleshooting & Support**
    *   Common Issues and contacting support

---

## 1. Introduction to Lyceum

### What is Lyceum?
Lyceum is an advanced, AI-powered personal learning platform designed to help you master any subject. It combines robust test generation capabilities (TestGen Suite) with structured courses, AI-assisted study tools, a collaborative chat environment, and personalized progress tracking. Whether you're preparing for exams, looking to deepen your understanding of complex topics, or exploring new areas of knowledge, Lyceum provides the tools and resources to enhance your learning journey. The platform features a professional, academic aesthetic with a touch of historic and gamified elements to make learning engaging and effective.

### Who is this Guide For?
This guide is intended for all users of the Lyceum platform, including:
*   **Students/Learners:** To understand how to create an account, navigate the platform, use the TestGen suite, enroll in and complete courses, utilize AI study aids, manage their profile, and interact with community features.
*   **Administrators:** To learn about the specialized Admin Panel for managing users, course content, platform settings, and other administrative tasks. Sections specific to administrators are clearly marked.

### Key Features Overview
*   **TestGen Suite:** Create customized tests from subject-specific question banks or your own materials. Track your performance and mastery over different chapters.
*   **Courses Platform:** Enroll in structured courses with lessons, video lectures, PDF materials, transcriptions, assignments, and exams.
*   **AI-Powered Tools:** Get AI explanations for questions, study material snippets, PDF content, and AI-generated formula sheets and chapter summaries. Improve your notes with AI assistance.
*   **AI Chat Studio:** Engage in focused conversations with the Lyceum AI (Lyra) on various topics or for general assistance.
*   **Personalized Dashboards:** Track your progress in TestGen and enrolled courses with detailed statistics and charts.
*   **Notes & Documents:** Create, manage, and (experimentally) share notes and documents related to your courses.
*   **Global Chat (Experimental):** Connect with other Lyceum users for discussion and help.
*   **Leaderboard & Credits:** Earn credits for platform engagement and see how you rank among other learners.
*   **Music & Sounds Player (Experimental):** Enhance your study environment with background music, ambient sounds, or binaural beats.
*   **Customizable Interface:** Personalize your Lyceum experience with theme selection and background customization.
*   **Admin Panel (Admin Only):** Comprehensive tools for platform management.

---

## 2. Getting Started

### Accessing Lyceum
Lyceum is a web-based application. You can access it using a modern web browser (like Chrome, Firefox, Edge, Safari) by navigating to its designated URL (e.g., `your-lyceum-domain.com` or `localhost` if running locally).

### Creating Your Account (Sign Up)
If you are a new user, you'll need to create an account.
1.  Navigate to the Lyceum login page. By default, the "Login" tab is active.
2.  Click the "Sign Up" button/tab to switch to the registration form.
3.  Fill in the required information:
    *   **Username:** Choose a unique username (3-20 characters, letters, numbers, or underscores only). This will be your public identifier (e.g., for mentions in chat).
    *   **Email:** Enter a valid email address. This will be used for account verification, password resets, and important communications.
    *   **Password:** Create a strong password (minimum 6 characters).
4.  Click the "Sign Up" button.
5.  Upon successful registration, you may be automatically logged in, or you might need to verify your email address depending on the platform's configuration. You will then be guided through the Initial Onboarding process.

### Logging In
If you already have an account:
1.  Navigate to the Lyceum login page. Ensure the "Login" tab is active.
2.  Enter your registered **Email or Username** in the first field.
3.  Enter your **Password** in the second field.
4.  Click the "Login" button.
5.  **Sign In with Google:** Alternatively, you can click the "Sign In with Google" button if you prefer to use your Google account. This will redirect you to Google for authentication and then back to Lyceum. If it's your first time using Google Sign-In with Lyceum, a Lyceum account linked to your Google email will be created.

### Password Reset
If you've forgotten your password:
1.  On the Login form, click the "**Forgot Password?**" link.
2.  A prompt will appear asking for the email address associated with your account.
3.  Enter your email address and click "OK" or "Submit".
4.  A password reset link will be sent to your email address. Follow the instructions in the email to create a new password.

### Initial Onboarding (First-time Setup)
After your first successful login (either via sign-up or first Google Sign-In), you will be presented with the Onboarding screen.
1.  **Welcome Message:** You'll see a welcome message and an explanation of the setup step.
2.  **Select Initial Subject (for TestGen):**
    *   A list of available global subjects for the TestGen suite will be displayed (e.g., "Fundamentals of Physics," "Artin Abstract Algebra").
    *   Click on the subject you primarily want to use for generating practice tests. This will be set as your initial active subject.
    *   **Note:** You can change this active subject and manage other subjects later via "Manage Subjects" in the sidebar.
3.  **OR Browse Courses:**
    *   Alternatively, you can click the "Browse Courses" button to skip selecting an initial TestGen subject and go directly to the course catalog.
4.  **Completion:**
    *   If you select a TestGen subject, the system will finalize your setup, and you'll be taken to the Home Dashboard.
    *   If you choose to browse courses, you'll be taken to the course listing page.

This onboarding step helps tailor your initial Lyceum experience.

---

## 3. Navigating Lyceum: The Main Interface

Once you've logged in and completed any initial onboarding, you'll be presented with the main Lyceum application interface. Understanding its layout is key to using the platform effectively.

### The Sidebar (Navigation Panel)
Located on the left side of the screen (on desktop views), the sidebar is your primary navigation tool. It dynamically changes based on whether you are in the general platform view or within a specific course.

#### Standard Navigation (`sidebar-standard-nav`)
This is the default sidebar view when you are not actively inside a specific course. It provides access to global features and dashboards.
*   **Home:** Takes you to your main Home Dashboard.
*   **TestGen Dashboard (Dropdown Menu):**
    *   **Generate Test:** Opens the interface to create new practice tests.
    *   **Studied Chapters:** Allows you to mark which chapters you have studied.
    *   **Exams Dashboard (TestGen):** View history and manage results for your TestGen exams.
    *   **Progress (TestGen):** See detailed performance charts for your active subject.
    *   **Manage Subjects:** Select your active TestGen subject or add/edit subjects.
    *   **Import / Export:** Manage your TestGen data backups.
*   **Browse Courses:** Takes you to the catalog of all available courses.
*   **My Courses:** Displays a list of all courses you are currently enrolled in.
*   **Music & Sounds (Experimental):** Opens the Music & Sounds Center.
*   **Leaderboard:** Shows the platform-wide leaderboard.
*   **Marketplace (Coming Soon / Experimental):** A planned area for using earned credits.
*   **(Divider)**
*   **Your Profile:** Access and manage your user profile.
*   **Inbox:** View messages from administrators or system notifications.
*   **Settings:** Configure your account, appearance, and other features.
*   **Global Chat (Experimental):** Opens the global chat modal.
*   **AI Chat Studio:** Access the dedicated interface for conversational AI.
*   **Admin Panel (Admin Only):** Provides access to platform management tools.
*   **Log Out:** Signs you out of your Lyceum account.

> **Sidebar Dropdowns:** Some items are dropdown menus. Clicking them reveals sub-links. An arrow icon indicates a dropdown. If a child link is active, the parent dropdown is also highlighted.
>
> **Active Link:** The currently active page is highlighted in the sidebar.

#### Course-Specific Navigation (`sidebar-course-nav`)
When you enter a specific course, the sidebar changes to provide relevant navigation:
*   **Course Dashboard:** Main overview page for the current course.
*   **Next Lesson:** A quick link to your next target chapter.
*   **Study Material:** Opens the full content menu (chapter list) for the course.
*   **Notes & Documents:** Access your personal notes for each chapter.
*   **Assignments & Exams:** View and access all assessments for the course.
*   **Course Progress:** View your performance, grades, and certificate eligibility.
*   **(Divider)**
*   **Back to My Courses:** Exits the current course view.

#### Theme Toggle (Bottom of Sidebar)
*   Click the sun/moon icon to toggle between **Light Mode** and **Dark Mode**.
*   Your preference is saved locally in your browser.

### Main Content Area (`main-content`)
This is the largest part of the screen, to the right of the sidebar. It's where content for the selected page is displayed (e.g., your dashboard, an online test, course materials, etc.). Content is often displayed within "Content Cards" with a frosted glass effect for better organization.

### Header Bar & User Section (Top of Main Content Area)
*   **Mobile Menu Button (Hamburger Icon):** On smaller screens, this button in the top-left toggles the sidebar's visibility. A red dot (ping) may appear for unread inbox messages.
*   **Lyceum Title (Mobile):** On mobile, the "Lyceum" title is displayed centrally when the sidebar is closed.
*   **User Section (Desktop):** On larger screens, your user information is typically in the top-right. This includes:
    *   Your profile picture (avatar).
    *   Your display name.
    *   An admin icon if you have administrator privileges.
*   **Subject Info (Desktop, for TestGen):** When the TestGen suite is active, this area shows the name of your currently selected TestGen subject.

### Mobile Navigation
*   On smaller screens, the sidebar is initially hidden. Use the **Mobile Menu Button** to open it.
*   When the sidebar is open on mobile, an **overlay** covers the main content area. Clicking the overlay closes the sidebar.

### Background & Opacity Customization (via `Settings` -> `Appearance`)
*   **App Background:** Change the overall background of the application.
    *   Choose from predefined images.
    *   Upload your own custom image (up to 3MB).
    *   Select a solid background color.
    *   Reset to the default.
*   **Card Opacity:** Adjust the transparency of "Content Cards".
    *   A slider sets opacity from slightly transparent (e.g., 10%) to fully opaque (100%).
    *   Lower opacity creates a "frosted glass" effect.

---

## 4. Home Dashboard

The Home Dashboard is your central landing page, providing a quick overview of your current activities.

### Overview
*   **Welcome Message:** Greets you by your display name.
*   The dashboard provides a snapshot of your progress in both the TestGen suite and the Courses platform.

### Test Generation Snapshot
This section summarizes your TestGen activities.
*   **Active Subject:** Displays the name of your active TestGen subject. If none is active, it prompts you to manage subjects.
*   **Key Statistics (for the active subject):**
    *   **Chapters Loaded:** Total number of chapters defined for the subject.
    *   **Chapters Marked as Studied:** Number of chapters you have marked as studied.
    *   **Exams Pending:** Number of PDF exams generated but not yet scored.
    *   **Recent Exam:** Details of your last completed exam (date and score).
*   **Quick Action Button:** A button like "Go to TestGen Dashboard" provides direct access.

### Courses Snapshot
This section summarizes your engagement with structured courses.
*   **Enrollment Status:** Displays the number of courses you are currently enrolled in. If none, it prompts you to "Browse Courses."
*   **Today's Focus (Full Enrollment Mode):** Highlights a suggested objective for the day based on your pace and progress.
*   **Next Task Button (Full Enrollment Mode):** If a specific "next task" is identified, a button appears to take you directly to that activity.
*   **Quick Action Button:** A button like "Go to My Courses" provides direct access to your list of enrolled courses.

---

## 5. TestGen Suite: Personalized Test Generation

### TestGen Dashboard
This is the main hub for the TestGen suite, accessed from the Home Dashboard or the sidebar (`TestGen Dashboard` -> `Generate Test`).

### Managing Subjects (for TestGen)
The TestGen suite operates on "subjects". Access this via `TestGen Dashboard` -> `Manage Subjects`.

*   **Functionality:**
    *   **List Available Subjects:** Displays all global subjects. Admins may see pending/rejected subjects.
    *   **Select Active Subject:** Click the "Select" button to make a subject your current active one.
    *   **View Subject Details:** Shows subject name, and for admins, creator info and status.
    *   **Add New Global Subject / Suggest Subject:**
        *   Admins can directly add new subjects.
        *   Non-admins can "suggest" a new subject for admin review.
    *   **Edit Global Subject Definition (Admin/Creator of Pending):** Admins can edit subject details.
    *   **Approve/Reject/Delete Global Subjects (Admin):** Admins manage the lifecycle of subjects.

#### Understanding Subject Details
*   **Name:** e.g., "Artin Abstract Algebra."
*   **Resource Directory Name:** The folder name on the server where files are stored.
*   **MCQ Markdown Filename:** The `.md` file containing multiple-choice questions.
*   **Optional Problem Filenames:** Similar files for other question types.
*   **Default Test Parameters:** Default settings for tests generated from this subject.

### Managing Studied Chapters
This feature lets you mark which chapters you've studied. Access via `TestGen Dashboard` -> `Studied Chapters`.
*   **Interface:** A list of all chapters for your active subject, each with a checkbox.
*   **Functionality:** Check the box to mark a chapter as "studied." Changes are saved automatically.
*   **Impact:** The "Test Studied Chapters Only" option will use these selections.

### Generating a New Test
Access this via `TestGen Dashboard` -> `Generate Test`.
1.  **Choose Scope:**
    *   **Test Studied Chapters Only:** Uses chapters you've marked in "Manage Studied Chapters."
    *   **Test Specific Chapters:** Lets you select chapters for this specific test.
2.  **Configure Your Test:**
    *   **Target Total Questions in Test:** Specify the desired number of questions.
    *   **Content from Main/Lecture Files:** Specify the number of MCQs and/or Problems from different source files.
    *   **Understanding "0 for auto":** Setting a count to 0 allows the system to auto-allocate questions from that source to meet the target total.
    *   **Test Timing:**
        *   **Subject Default:** Uses the subject's default duration.
        *   **Calculated:** System estimates a duration.
        *   **Custom Duration:** Enter a specific duration in minutes.
3.  **Choose Format & Start:**
    *   **Start Online Test:** Immediately launches the online test interface.
    *   **Generate PDF Test:** Generates and downloads a "Questions PDF" and a "Solutions PDF." The exam is added to your pending list for scoring.
    *   **Generate LaTeX Source (.tex):** Generates and downloads `.tex` source files for questions and solutions.

### Taking an Online Test (TestGen)
*   **The Test Interface:**
    *   **Header:** Displays Exam ID and a countdown Timer.
    *   **Question Area:** Shows the current question text and image.
    *   **Answer Area:**
        *   For MCQs: Radio buttons for options.
        *   For Problems: A rich text input area with a **LaTeX Toolbar** and a live preview.
    *   **Navigation Footer:** "Previous" and "Next" buttons, question counter, and a "Submit Test" button.
*   **Answering MCQs:** Click the radio button for your choice. It's saved automatically.
*   **Answering Problems (LaTeX Input):** Use the toolbar or type raw LaTeX directly into the input area. Use `$...$` for inline math and `$$...$$` for display math.
*   **Timer and Submission:** The test submits automatically when the timer runs out. You can also submit early.
*   **After Submission:** The system grades your test and displays results, scores, and AI feedback.

### TestGen Exams Dashboard & History
Access via `TestGen Dashboard` -> `Exams Dashboard`.
*   **Pending PDF Exams:**
    *   Lists PDF exams you've generated but haven't scored.
    *   **Enter Results:** Input how many questions you got wrong per chapter.
    *   **Delete Pending:** Removes the pending entry.
*   **Completed Exam History (New Centralized Store):**
    *   Shows all completed exams (online and scored PDFs).
    *   Displays Exam ID, type, date, and score.
    *   **Review:** Opens the detailed Exam Review UI.
    *   **Delete:** Permanently deletes the exam history record.

#### Reviewing Exam Details & AI Feedback (TestGen Exams)
When you click "Review" for a completed exam:
*   **Overall Summary:** Your score, percentage, pass/fail status, etc.
*   **Overall AI Feedback:** Strengths, weaknesses, and study recommendations.
*   **Question Breakdown:** A list of all questions.
    *   **Question Display:** The question, image, and options.
    *   **Your Answer & Correct Answer:** Clearly indicated.
    *   **Score Awarded:** For that specific question.
    *   **AI Marking Feedback (for Problems):** Detailed feedback on your problem-solving steps.
    *   **AI Explanation (Button):** Click to get a step-by-step explanation of the concepts. This opens an interactive chat-like panel.
    *   **Report Issue (Button):** Report problems with the question or AI feedback.

### TestGen Progress Dashboard
Access via `TestGen Dashboard` -> `Progress (TestGen)`.
*   **Difficulty Score Chart:** Shows the AI-calculated difficulty score for each chapter based on your performance.
*   **Consecutive Mastery Chart:** Shows how many consecutive tests you've "mastered" each chapter on.
*   **Questions Attempted Chart:** Shows the total number of questions you've attempted from each chapter.
*   **Wrong Answers Chart:** Shows the total number of questions you've answered incorrectly from each chapter.

---

## 6. Courses Platform

The Courses Platform provides structured learning experiences.

### Browsing Courses
Access by clicking "Browse Courses" in the sidebar.
*   **Interface:**
    *   **Search Bar:** Search for courses by name, tag, or keyword.
    *   **"Add Course" / "Suggest Course" Button:** Admins can add courses directly; users can suggest new courses.
    *   **Course List:** Displays available courses as cards, showing thumbnail, name, description, tags, and enrollment status.

### Viewing Course Details
Clicking on a course takes you to its detail page.
*   **Information Displayed:** Full name, description, creator, tags, chapter list, prerequisites, corequisites, etc.
*   **Actions:**
    *   **Enroll Button:** Enroll in an approved course.
    *   **Go to Course Dashboard Button:** If already enrolled.
    *   **Report Course Button:** Report issues with an approved course.
    *   **Admin Actions (Admin Only):** Edit, approve, reject, or delete courses.

### Enrolling in a Course & Setting Your Pace
When you click "Enroll":
*   **Enrollment Mode:**
    *   **Full Enrollment:** Standard mode with access to all activities, progress tracking, and certificate eligibility.
    *   **Viewer Mode:** Access to study materials and notes only. Assignments, exams, and progress tracking are disabled.
*   **Learning Pace (for Full Enrollment only):**
    *   Choose a pace (Compact, Mediocre, Lenient, or Custom) to set your daily learning targets.
*   **Confirmation:** Click "Enroll & Start Learning" to finalize.

### My Courses Dashboard
Access by clicking "My Courses" in the sidebar.
*   Displays a list of all courses you are enrolled in (in Full or Viewer mode).
*   Each card shows course name, status, grade, and progress.

### Course-Specific Dashboard
This is your main hub for an individual course.
*   **Cover Image & Course Title:** The course's branding.
*   **Unenroll Button:** Allows you to leave the course (progress will be deleted).
*   **Today's Objective (Full Enrollment Mode):** Suggests what you should focus on today.
*   **Quick Links (Cards/Buttons):**
    *   **Study Material (All Chapters):** Go to the course content menu.
    *   **Next Lesson:** A quick link to your next chapter.
    *   **Assignments & Exams:** View all available assessments.
    *   **Notes & Documents:** Access your notes for the course.
    *   **Detailed Progress:** View charts, score breakdowns, and your certificate.

### Study Material
*   **Chapter Navigation:** "Previous" and "Next" buttons to navigate between parts of a chapter.
*   **Video Player (YouTube):** Embedded player for lecture videos with progress tracking.
*   **Video Transcription:** If available, the video's transcript is displayed below the player, with interactive highlighting and click-to-seek functionality.
*   **PDF Viewer:** Embedded viewer for PDF documents, with page navigation and progress tracking.
*   **AI Tools for Study Material:**
    *   **Explain Selection (Transcription):** Select text in the transcript and ask the AI for an explanation.
    *   **Ask AI (About Transcription/PDF Page/Full PDF Document):** Ask the AI questions about the content of the video or PDF.
*   **Chapter Formula Sheet (AI Generated):** Generate a concise formula sheet for the current chapter.
*   **Chapter Summary (AI Generated):** Generate a comprehensive summary of the chapter.
*   **Skip Exam (Testing out of a Chapter - Full Enrollment Mode):** If a chapter is not yet "studied," you can take a "Skip Exam." Passing it marks the chapter as studied.

### Assignments & Exams (Course-Specific)
*   **Types of Activities:** Daily Assignments, Weekly Exams, Midcourse Exams, Final Exams. These become available as you progress.
*   **Interface:** Lists activities by type, showing name/ID and status (Available, Completed, Locked).
*   **Taking Activities:** Clicking "Start" launches an online test interface similar to TestGen.

### Course Progress Details
*   **Overall Grade & Performance:** Displays your total mark, letter grade, attendance score, and current pace.
*   **Performance Charts:** Line and bar charts showing your scores over time for assignments and exams.
*   **Score Breakdown Table:** Itemizes how different components contribute to your total mark.
*   **Course Certificate (If Eligible):** If your course is "Completed" with a passing grade, a certificate preview appears. You can regenerate the preview and download it as a PNG or PDF.

### Notes & Documents (Per Course, Per Chapter)
*   **Chapter List Menu:** Select a chapter to view its dedicated notes panel.
*   **Notes Panel (for a specific chapter):**
    *   **Your Notes Section:** A list of all notes and files you've created or uploaded for the chapter. Actions include View, Edit, Delete, and AI Tools.
    *   **AI Tools (for Notes):** Review a note for accuracy, or use OCR & LaTeX conversion for images.
    *   **Add Note Button:** Create a new text-based note.
    *   **Upload File Button:** Upload files (PDF, TXT, JPG, etc.) to associate with the chapter.
    *   **Shared Notes Section (Future/Experimental):** A planned area to view notes shared by other users.

---

## 7. AI Chat Studio

This is a dedicated interface for free-form conversation with Lyra, the Lyceum AI.

*   **Interface Overview:**
    *   **Sessions Sidebar (left):**
        *   **New Chat Button:** Start a new conversation.
        *   **Chat History List:** A list of your previous chat sessions. Click to load a session.
        *   **Delete Session Button:** Appears on hover to delete a session.
    *   **Main Chat Area (right):**
        *   **Chat Header:** Displays the current session's name.
        *   **Messages Container:** Where your conversation with Lyra unfolds. User messages and AI responses are displayed chronologically.
        *   **Input Area (bottom):** A multi-line textarea to type your messages and a "Send" button.
*   **Starting a New Chat:** Click the "New Chat" button, optionally give it a name, and start typing.
*   **Sending Messages & Using LaTeX:**
    *   Type your question in the textarea.
    *   Use `$x^2$` for inline math and `$$ \int_a^b f(x) dx $$` for display math.
    *   Press Enter or click the "Send" button.
*   **Managing Chat Sessions:**
    *   **Switching:** Click any session in the Chat History list to load it.
    *   **Deleting:** Hover over a session and click the delete icon.

---

## 8. Community & Engagement

### Leaderboard
*   **Accessing:** Click "Leaderboard" in the main sidebar.
*   **Purpose:** Displays a ranked list of top users based on accumulated "Credits."
*   **Earning Credits:** Credits are awarded for activities like completing exams, enrolling in courses, and other forms of engagement.

### Global Chat (Experimental)
> **Note:** This feature is experimental. Availability depends on your settings.

*   **Accessing:** Click "Global Chat" in the sidebar to open a full-screen chat overlay.
*   **Functionality:**
    *   **Sending Messages:** Type your message and send. It's visible to all users in the chat.
    *   **Mentions:** Use `@username` to notify a specific user or `@everyone` to notify all active users.
    *   **Replying:** Click the reply icon on a message to reply in context.
    *   **Pinning Messages (Admin Only):** Admins can pin important messages.
    *   **Deleting Messages:** You can delete your own messages. Admins can delete any message.

### Marketplace (Coming Soon / Experimental)
*   A planned feature where users can spend credits earned through platform engagement. Currently a placeholder.

---

## 9. Your Profile & Settings

### Viewing Your Profile
*   **Accessing:** Click "Your Profile" in the main sidebar.
*   **Information Displayed:** Avatar, Display Name, Username, Email, Credits, UID, and Completed Course Badges.
*   **Editable Fields:** You can change your Profile Picture and Display Name directly on this page.

### Editing Your Profile (Display Name, Avatar via Cropper)
*   **Changing Display Name:** Type your new name in the input field and save.
*   **Changing Profile Picture (Avatar):**
    1.  Click on your avatar to upload a new image.
    2.  Use the cropper tool to select the desired portion of the image.
    3.  Click "Crop & Use" to save the new avatar.

### Settings Panel
Access by clicking "Settings" in the sidebar. It's organized into tabs.

*   **A. Profile & Account Management Tab:**
    *   **Display Name & Avatar:** Update your name and picture.
    *   **Account Management:**
        *   **Username:** Change your unique `@username`.
        *   **Email:** Change your registered email address (requires verification).
        *   **Change Password:** Update your password.
    *   **Danger Zone:**
        *   **Log Out Button:** Logs you out of your account.
        *   **Delete My Account Button:** **CRITICAL:** Initiates the permanent deletion of your account and all associated data. This action is irreversible and requires multiple confirmations.

*   **B. Appearance Settings Tab:**
    *   Customize the application's **background** (image or solid color) and the **Card Opacity** (transparency of UI elements).

*   **C. Experimental Features Tab:**
    *   A list of beta features (e.g., Global Chat, Marketplace, Music Player) that you can toggle on or off.

*   **D. AI Model & System Prompt Settings (Advanced Users / Admin-restricted):**
    *   **Model Configuration:** Select your preferred primary and fallback AI models.
    *   **System Prompts:** View and edit the underlying instructions (prompts) given to the AI for specific tasks. You can set your own custom prompts, which override platform defaults.

---

## 10. Music & Sounds Player (Experimental)

> **Note:** This feature is experimental. Availability depends on your settings.

Enhance your study environment with background music and ambient sounds.

*   **Accessing:** Click "Music & Sounds" in the sidebar.
*   **Interface:** A two-column layout.
    *   **Left Column (Sources & Libraries):** Select from music sources (`YouTube Music`), `Ambient Sounds`, `Study Music` (Lo-fi, Chill), `Binaural Beats`, or `Saved Playlists`.
    *   **Right Column (Main Player & Playlist):** The main playback area, controls, and current playlist.
*   **Main Player Controls:** Standard controls like Play/Pause, Next/Previous, Shuffle, Repeat, Progress Slider, and Volume Slider.
*   **Playlist Management:** View the current playlist. Save the current playlist for later or load a previously saved one.
*   **Mini Player (Floatie):** A small, persistent music player that floats over the interface, allowing you to control music from anywhere on the platform.
*   **Sound Settings:** Toggle UI sound effects and control ambient sound volume.

---

## 11. Import / Export Test Data

This feature allows you to back up and restore your **TestGen data only**. It does not affect data from the Courses Platform.

*   **Accessing:** Click `Import / Export` under the `TestGen Dashboard` dropdown.
*   **Exporting Your TestGen Data:**
    1.  Click the "Export My Test Data" button.
    2.  The system gathers your TestGen subjects, chapter progress, studied status, and pending exams into a JSON file.
    3.  The file (`test_generator_backup_...json`) is automatically downloaded. Save it in a safe place.
*   **Importing TestGen Data:**
    > **WARNING:** This is a destructive operation that will **completely overwrite** your current TestGen data.
    1.  Click the "Import Test Data from File..." button.
    2.  Select the JSON backup file you previously exported.
    3.  Confirm the overwrite in the warning dialog.
    4.  The system validates the file, re-syncs chapter data with the latest server versions while preserving your progress, and replaces your old TestGen data with the data from the file.

---

## 12. Admin Panel (For Administrators Only)

> **NOTE:** This section is only for users with administrator privileges.

*   **Accessing:** Click the "Admin Panel" link in the main sidebar.
*   **Sections:** The admin panel has its own sidebar for navigating different administrative areas: Overview, Users, Courses, Moderation, Music & Sounds, Database, System, and Testing Aids.

*   **Key Functionalities:**
    *   **Overview:** View platform statistics and send global announcements to all users.
    *   **User Management:** Search for users, view their details, manage their course progress, toggle their admin status (Primary Admin only), and change usernames.
    *   **Global Course Management:** Approve/reject/delete courses, add new global courses, and edit existing course definitions.
    *   **Moderation:** View and respond to user feedback and exam issue reports.
    *   **Music & Sound Library Management:** Edit the predefined lists of music and sounds available to all users.
    *   **Database Management (Advanced Admin):** **DANGER:** A direct browser for the Firestore database. Use with extreme caution.
    *   **System Operations (Primary Admin):** Manage admin tasks, global AI system prompts, and other system-level settings.
    *   **Admin Testing Aids:** **DANGER:** Tools to quickly alter a user's progress for testing purposes. Use with extreme care.

---

## 13. Troubleshooting & Support

### A. Common Issues & Solutions

1.  **Login Problems**
    *   **Issue:** Incorrect email/username or password.
        *   **Solution:** Double-check spelling, ensure Caps Lock is off, and try the "Forgot Password?" link.
    *   **Issue:** Google Sign-In popup closes or shows an error.
        *   **Solution:** Ensure your browser isn't blocking pop-ups. Check if third-party cookies are required and temporarily enable them.

2.  **Display & Rendering Issues**
    *   **Issue:** Mathematical formulas (LaTeX) are not rendering correctly.
        *   **Solution:** Wait a few seconds for MathJax to load. Try a hard refresh (Ctrl+Shift+R or Cmd+Shift+R). Check for interfering browser extensions.
    *   **Issue:** Page layout looks broken or styles are missing.
        *   **Solution:** Try a hard refresh. Ensure your browser zoom is at 100%. As a last resort, clear your browser's cache and cookies for the site.

3.  **TestGen Suite Issues**
    *   **Issue:** "No chapters loaded" or "No questions found."
        *   **Solution (Admin):** Ensure the subject's definition file (.md) exists at the correct path on the server and is formatted correctly.
    *   **Issue:** Test generation fails or generates fewer questions than requested.
        *   **Solution:** There may not be enough available questions in the selected chapters. Try selecting more chapters or reducing the total question count.
    *   **Issue:** PDF Test download fails or the PDF is blank.
        *   **Solution:** This could be a server-side issue with the PDF generation service or an issue with overly complex content. Try a smaller test. Ensure pop-ups are not blocked.

4.  **Courses Platform Issues**
    *   **Issue:** Cannot enroll in a course.
        *   **Solution:** Only "approved" courses can be enrolled in. Check "My Courses" to see if you are already enrolled.
    *   **Issue:** Video player shows an error.
        *   **Solution:** The video may have been removed or made private on YouTube. Try a standard network/browser refresh.
    *   **Issue:** Progress (video watch time, PDF page) not saving.
        *   **Solution:** Ensure a stable internet connection. If you enrolled in "Viewer Mode," progress is not tracked.

5.  **AI Features (Chat Studio, Explanations, etc.)**
    *   **Issue:** AI response is "Error generating..." or no response is given.
        *   **Solution:** The AI service may be temporarily unavailable. Your prompt might be too long or complex. Try simplifying your query.
    *   **Issue:** AI explanations seem inaccurate.
        *   **Solution:** AI is a tool and not infallible. Use the "Report Issue" button or contact support. Always cross-reference critical information.

### B. Reporting Issues
*   **Specific Question/Content Issues:** When reviewing an exam, use the **"Report Issue"** button available for each question.
*   **General Feedback/Bugs:** Use the "Contact Admin" feature or any dedicated "Feedback" link on the platform. Be as specific as possible: what you were doing, steps to reproduce, what you expected, and what actually happened.

### C. Contacting Support
*   If you cannot resolve an issue, look for a "Support," "Help," or "Contact Us" link, often in the page footer or profile menu.
*   Provide as much detail as possible in your request.

### D. Tips for a Smooth Experience
1.  **Use a Modern Browser:** Use up-to-date versions of Chrome, Firefox, Edge, or Safari.
2.  **Stable Internet Connection:** Many features rely on server communication.
3.  **Keep Data Backups (TestGen):** Regularly use the "Export Test Data" feature.
4.  **Be Mindful of Experimental Features:** They may have bugs. Use them with this understanding.
5.  **Provide Constructive Feedback:** Your feedback is valuable for improving Lyceum.