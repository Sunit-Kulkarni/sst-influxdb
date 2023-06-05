import Image from 'next/image'
import styles from './page.module.css'

export default function Home() {
  return (
    <main className={styles.main}>

      <div className={styles.center}>
        <div className={styles.grid}>
          <a
            href="https://nextjs.org/docs?utm_source=create-next-app&utm_medium=appdir-template&utm_campaign=create-next-app"
            className={styles.card}
            target="_blank"
            rel="noopener noreferrer"
          >
            <Image
              className={styles.logo}
              src={"/VinoBarrelCabernet.png"}
              width={208}
              height={928}
            />
            
          </a>

          <a
            href="https://nextjs.org/learn?utm_source=create-next-app&utm_medium=appdir-template&utm_campaign=create-next-app"
            className={styles.card}
            target="_blank"
            rel="noopener noreferrer"
          >
            <Image
              className={styles.logo}
              src={"/VinoBarrelZinfandel.png"}
              width={208}
              height={928}
            />
            
          </a>
        </div>
      </div>

    </main>
  )
}
