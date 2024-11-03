/** @type {import('ts-jest').JestConfigWithTsJest} **/
const config = {
  testEnvironment: "node",
  transform: {
    "^.+.tsx?$": ["ts-jest", {}],
  },
  moduleNameMapper: {
    '^@/(.*)$': '<rootDir>/typescript/$1',
  },
}

export default config
