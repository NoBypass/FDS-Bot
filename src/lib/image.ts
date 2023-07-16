import { chromium, Browser, Page } from 'playwright'
import { Measurements } from '../types/common'

const generateImageFromHtml = async (
  html: string,
  css: string,
  measurements?: Measurements,
) => {
  const browser: Browser = await chromium.launch()
  const context = await browser.newContext()
  const page: Page = await context.newPage()

  const pageContent = `
  <html lang="en">
    <head>
      <title>.</title>
      <style>
        @import url('https://fonts.googleapis.com/css2?family=Raleway&display=swap');
        * {
            font-family: 'Raleway', sans-serif;
            color: #ffffff;
            background-color: #57F287;
        }
        ${css}
      </style>
    </head>
    <body>
      ${html}
    </body>
  </html>
  `
  await page.setContent(pageContent)
  if (measurements) await page.setViewportSize(measurements)

  const sc = await page.screenshot()
  await browser.close()
  return sc
}

export default generateImageFromHtml
