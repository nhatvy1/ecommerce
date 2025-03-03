export default () => ({
  port: parseInt(process.env.PORT) || 5000,
  database: process.env.MONGO_URI
})
