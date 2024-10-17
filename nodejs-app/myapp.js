const express = require('express');
const app = express();

// environment variables with defaults
const PORT = process.env.PORT || 3000;
const GREETING = process.env.GREETING || 'Hello';
const NAME = process.env.NAME || 'Docker';

app.get('/', (req, res) => {
  res.send(`${GREETING}, ${NAME}! (from Node.js)`);
});

app.listen(PORT, () => {
  console.log(`Server running on port ${PORT}`);
});
