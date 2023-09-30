# KTUNotifier

KTUNotifier is a real-time notification feed that provides updates from ktu.edu.live and pushes them to notification channels like Telegram. It allows users to stay informed about the latest announcements, news, and updates from the Kerala Technological University (KTU) website.

## Features

- **Real-time Updates**: KTUNotifier fetches updates from ktu.edu.live in real-time, ensuring that users receive the latest information as soon as it is published on the website.
- **Notification Channels**: KTUNotifier integrates with popular messaging platforms like Telegram, allowing users to receive notifications directly on their devices.
- **Automatic Updates**: The application automatically fetches updates at regular intervals, ensuring that users are always up to date with the latest information from ktu.edu.live.

## Environment Variables

KTUNotifier requires the following environment variables to be set:

- `BOT_TOKEN`: The Telegram bot token used to send notifications to users.
- `CHAT_ID`: The Telegram chat ID used to send notifications to users.

## Installation

To use KTUNotifier, follow these steps:

```bash
# Clone the repository
git clone https://github.com/AmarnathCJD/KTUNotifier.git

# Change the working directory
cd KTUNotifier

# Install and compile
go build .
```

## Usage

Once KTUNotifier is set up and running, it will automatically fetch updates from ktu.edu.live and push them to the configured notification channels, such as Telegram. Users can customize their notification settings, filters, and channels through the provided user interface.

## Contributing

Contributions to KTUNotifier are welcome! If you would like to contribute, please follow these guidelines:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Implement your changes and ensure they are working correctly.
4. Commit your changes and push them to your forked repository.
5. Submit a pull request, explaining the changes you have made and their purpose.

## License

KTUNotifier is released under the [GNU General Public License v3.0](LICENSE). Feel free to modify and distribute the application according to the terms of the license.

Please note that the usage of KTUNotifier is subject to the terms and conditions set by ktu.edu.live. Ensure compliance with their terms of service when using KTUNotifier.