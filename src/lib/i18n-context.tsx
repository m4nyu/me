"use client"

import React, { createContext, useContext, useState, useEffect, type ReactNode } from "react"

type Locale = "en" | "de" | "fr" | "es" | "it" | "pt" | "nl" | "ru" | "ja" | "ko" | "zh" | "ar"

interface I18nContextType {
  locale: Locale
  setLocale: (locale: Locale) => void
  t: (key: string) => string
}

const I18nContext = createContext<I18nContextType | null>(null)

const translations: Record<Locale, Record<string, string>> = {
  en: {
    // Navigation
    "nav.offer": "Offer",
    "nav.pricing": "Pricing", 
    "nav.qa": "Q&A",
    
    // Hero
    "hero.title": "Professional Cloud Browser Solutions",
    "hero.subtitle": "We provide cutting-edge cloud browser technology for businesses and developers. Scale your operations with our reliable, secure, and high-performance browser infrastructure.",
    "hero.cta": "Get in Touch",
    "hero.maskTitle": "Advanced Browser Infrastructure",
    "hero.maskSubtitle": "Experience the future of cloud computing with our revolutionary browser technology that adapts to your needs.",
    
    // Pricing
    "pricing.title": "Pricing",
    "pricing.subtitle": "Choose the plan that best fits your business needs",
    "pricing.mostPopular": "MOST POPULAR",
    "pricing.hourly.name": "Hourly",
    "pricing.hourly.price": "$100",
    "pricing.hourly.period": "/hour",
    "pricing.hourly.description": "Perfect for quick tasks and testing",
    "pricing.hourly.feature1": "Pay as you go",
    "pricing.hourly.feature2": "Instant setup",
    "pricing.hourly.feature3": "Basic support",
    "pricing.hourly.feature4": "No commitments",
    "pricing.daily.name": "Daily",
    "pricing.daily.price": "$800",
    "pricing.daily.period": "/day",
    "pricing.daily.description": "Ideal for ongoing projects",
    "pricing.daily.feature1": "Full day access",
    "pricing.daily.feature2": "Priority support",
    "pricing.daily.feature3": "Advanced features",
    "pricing.daily.feature4": "24/7 availability",
    "pricing.daily.feature5": "Custom configurations",
    "pricing.startup.name": "Startups",
    "pricing.startup.price": "Co-founder",
    "pricing.startup.period": "",
    "pricing.startup.description": "Equity participation and ownership stake",
    "pricing.startup.feature1": "Equity-based partnership",
    "pricing.startup.feature2": "Active participation",
    "pricing.startup.feature3": "Long-term commitment",
    "pricing.startup.feature4": "Shared ownership",
    "pricing.startup.feature5": "Strategic involvement",
    
    // Q&A
    "qa.title": "Questions",
    "qa.subtitle": "Everything you need to know about working together",
    "qa.services.question": "What types of services do I offer?",
    "qa.services.answer": "I specialize in full-stack web development, API design and implementation, cloud architecture, database optimization, and custom software solutions. Whether you need a new application built from scratch, legacy code modernization, or technical consulting, I can help bring your project to life.",
    "qa.approach.question": "How do I approach new projects?",
    "qa.approach.answer": "I begin with a thorough discovery phase to understand your requirements, followed by a detailed project proposal with clear milestones. I work in iterative sprints with regular check-ins, ensuring transparency and flexibility throughout the development process. Timelines are estimated based on project complexity and agreed upon before starting.",
    "qa.pricingInfo.question": "What is my pricing structure?",
    "qa.pricingInfo.answer": "I offer flexible pricing models including hourly rates for short-term work, daily rates for ongoing projects, and equity-based partnerships for startups. Each project is unique, and I provide customized quotes based on scope, complexity, and timeline. Contact me for a detailed estimate tailored to your specific needs.",
    "qa.technologies.question": "What technologies do I work with?",
    "qa.technologies.answer": "I'm proficient in modern web technologies including React, Next.js, Node.js, TypeScript, Python, and various databases. I also have experience with cloud platforms like AWS and GCP, containerization with Docker, and CI/CD pipelines. I stay current with industry trends and can adapt to your existing tech stack.",
    "qa.quality.question": "How do I ensure code quality?",
    "qa.quality.answer": "I follow industry best practices including test-driven development, code reviews, and comprehensive documentation. All code is version-controlled, well-commented, and built with scalability in mind. I maintain clear communication throughout the project and provide post-launch support to ensure smooth deployment.",
    
    // Footer
    "footer.imprint": "Imprint",
    "footer.gdpr": "GDPR",
    "footer.terms": "Terms of Service",
    "footer.brandName": "4nuel",
    
    // Language picker
    "language.choose": "Choose Language",
    
    // Mode
    "mode.dark": "Dark",
    "mode.light": "Light", 
    "mode.auto": "Auto",
  },
  de: {
    "nav.offer": "Angebot",
    "nav.pricing": "Preise",
    "nav.qa": "Fragen",
    "hero.title": "Professionelle Cloud-Browser-Lösungen",
    "hero.subtitle": "Wir bieten modernste Cloud-Browser-Technologie für Unternehmen und Entwickler. Skalieren Sie Ihre Abläufe mit unserer zuverlässigen, sicheren und leistungsstarken Browser-Infrastruktur.",
    "hero.cta": "Kontakt aufnehmen",
    "hero.maskTitle": "Fortschrittliche Browser-Infrastruktur",
    "hero.maskSubtitle": "Erleben Sie die Zukunft des Cloud-Computing mit unserer revolutionären Browser-Technologie, die sich an Ihre Bedürfnisse anpasst.",
    "pricing.title": "Preise",
    "pricing.subtitle": "Wählen Sie den Plan, der am besten zu Ihren Geschäftsanforderungen passt",
    "pricing.mostPopular": "AM BELIEBTESTEN",
    "pricing.hourly.name": "Stündlich",
    "pricing.hourly.price": "100€",
    "pricing.hourly.period": "/Stunde",
    "pricing.hourly.description": "Perfekt für schnelle Aufgaben und Tests",
    "pricing.hourly.feature1": "Nach Verbrauch zahlen",
    "pricing.hourly.feature2": "Sofortige Einrichtung",
    "pricing.hourly.feature3": "Basis-Support",
    "pricing.hourly.feature4": "Keine Verpflichtungen",
    "pricing.daily.name": "Täglich",
    "pricing.daily.price": "800€",
    "pricing.daily.period": "/Tag",
    "pricing.daily.description": "Ideal für laufende Projekte",
    "pricing.daily.feature1": "Ganztägiger Zugang",
    "pricing.daily.feature2": "Prioritätssupport",
    "pricing.daily.feature3": "Erweiterte Funktionen",
    "pricing.daily.feature4": "24/7 Verfügbarkeit",
    "pricing.daily.feature5": "Individuelle Konfigurationen",
    "pricing.startup.name": "Startups",
    "pricing.startup.price": "Mitgründer",
    "pricing.startup.period": "",
    "pricing.startup.description": "Eigenkapitalbeteiligung und Eigentumsanteil",
    "pricing.startup.feature1": "Eigenkapital-basierte Partnerschaft",
    "pricing.startup.feature2": "Aktive Beteiligung",
    "pricing.startup.feature3": "Langfristige Verpflichtung",
    "pricing.startup.feature4": "Geteiltes Eigentum",
    "pricing.startup.feature5": "Strategische Beteiligung",
    "qa.title": "Fragen",
    "qa.subtitle": "Alles was Sie über die Zusammenarbeit wissen müssen",
    "qa.services.question": "Welche Arten von Dienstleistungen biete ich an?",
    "qa.services.answer": "Ich spezialisiere mich auf Full-Stack-Webentwicklung, API-Design und -Implementation, Cloud-Architektur, Datenbankoptimierung und maßgeschneiderte Software-Lösungen. Ob Sie eine neue Anwendung von Grund auf benötigen, Legacy-Code-Modernisierung oder technische Beratung - ich kann Ihnen dabei helfen, Ihr Projekt zum Leben zu erwecken.",
    "qa.approach.question": "Wie gehe ich an neue Projekte heran?",
    "qa.approach.answer": "Ich beginne mit einer gründlichen Analysephase, um Ihre Anforderungen zu verstehen, gefolgt von einem detaillierten Projektvorschlag mit klaren Meilensteinen. Ich arbeite in iterativen Sprints mit regelmäßigen Check-ins und gewährleiste Transparenz und Flexibilität während des gesamten Entwicklungsprozesses.",
    "qa.pricingInfo.question": "Wie ist meine Preisstruktur?",
    "qa.pricingInfo.answer": "Ich biete flexible Preismodelle, einschließlich Stundensätze für kurzfristige Arbeiten, Tagessätze für laufende Projekte und eigenkapitalbasierte Partnerschaften für Startups. Jedes Projekt ist einzigartig, und ich erstelle maßgeschneiderte Angebote basierend auf Umfang, Komplexität und Zeitplan.",
    "qa.technologies.question": "Mit welchen Technologien arbeite ich?",
    "qa.technologies.answer": "Ich beherrsche moderne Web-Technologien wie React, Next.js, Node.js, TypeScript, Python und verschiedene Datenbanken. Ich habe auch Erfahrung mit Cloud-Plattformen wie AWS und GCP, Containerisierung mit Docker und CI/CD-Pipelines.",
    "qa.quality.question": "Wie gewährleiste ich Code-Qualität?",
    "qa.quality.answer": "Ich folge bewährten Praktiken der Branche, einschließlich testgetriebener Entwicklung, Code-Reviews und umfassender Dokumentation. Der gesamte Code ist versionskontrolliert, gut kommentiert und auf Skalierbarkeit ausgelegt.",
    "footer.imprint": "Impressum",
    "footer.gdpr": "DSGVO",
    "footer.terms": "Nutzungsbedingungen",
    "footer.brandName": "4nuel",
    "language.choose": "Sprache wählen",
    "mode.dark": "Dunkel",
    "mode.light": "Hell",
    "mode.auto": "Auto",
  },
  // Add more languages as needed - keeping it concise for now
  fr: {
    "nav.offer": "Offre",
    "nav.pricing": "Tarifs",
    "nav.qa": "Questions",
    "hero.title": "Solutions de Navigateur Cloud Professionnelles",
    "hero.subtitle": "Nous fournissons une technologie de navigateur cloud de pointe pour les entreprises et les développeurs.",
    "hero.cta": "Nous Contacter",
    "hero.maskTitle": "Infrastructure de Navigateur Avancée",
    "hero.maskSubtitle": "Découvrez l'avenir du cloud computing avec notre technologie révolutionnaire.",
    "pricing.title": "Tarifs",
    "pricing.subtitle": "Choisissez le plan qui correspond le mieux à vos besoins",
    "pricing.mostPopular": "LE PLUS POPULAIRE",
    "pricing.hourly.name": "Horaire",
    "pricing.hourly.price": "100€",
    "pricing.hourly.period": "/heure",
    "pricing.hourly.description": "Parfait pour les tâches rapides",
    "pricing.hourly.feature1": "Paiement à l'utilisation",
    "pricing.hourly.feature2": "Configuration instantanée",
    "pricing.hourly.feature3": "Support de base",
    "pricing.hourly.feature4": "Aucun engagement",
    "pricing.daily.name": "Quotidien",
    "pricing.daily.price": "800€",
    "pricing.daily.period": "/jour",
    "pricing.daily.description": "Idéal pour les projets en cours",
    "pricing.daily.feature1": "Accès toute la journée",
    "pricing.daily.feature2": "Support prioritaire",
    "pricing.daily.feature3": "Fonctionnalités avancées",
    "pricing.daily.feature4": "Disponibilité 24/7",
    "pricing.daily.feature5": "Configurations personnalisées",
    "pricing.startup.name": "Startups",
    "pricing.startup.price": "Co-fondateur",
    "pricing.startup.period": "",
    "pricing.startup.description": "Participation en capital",
    "pricing.startup.feature1": "Partenariat basé sur les capitaux propres",
    "pricing.startup.feature2": "Participation active",
    "pricing.startup.feature3": "Engagement à long terme",
    "pricing.startup.feature4": "Propriété partagée",
    "pricing.startup.feature5": "Implication stratégique",
    "qa.title": "Questions",
    "qa.subtitle": "Tout ce que vous devez savoir",
    "qa.services.question": "Quels types de services j'offre ?",
    "qa.services.answer": "Je me spécialise dans le développement web full-stack, la conception d'API, l'architecture cloud et les solutions logicielles personnalisées.",
    "qa.approach.question": "Comment j'aborde les nouveaux projets ?",
    "qa.approach.answer": "Je commence par une phase de découverte approfondie pour comprendre vos exigences, suivie d'une proposition détaillée avec des jalons clairs.",
    "qa.pricingInfo.question": "Quelle est ma structure tarifaire ?",
    "qa.pricingInfo.answer": "J'offre des modèles de tarification flexibles incluant des tarifs horaires, quotidiens et des partenariats basés sur les capitaux propres.",
    "qa.technologies.question": "Avec quelles technologies je travaille ?",
    "qa.technologies.answer": "Je maîtrise les technologies web modernes incluant React, Next.js, Node.js, TypeScript, Python et diverses bases de données.",
    "qa.quality.question": "Comment j'assure la qualité du code ?",
    "qa.quality.answer": "Je suis les meilleures pratiques incluant le développement dirigé par les tests, les revues de code et une documentation complète.",
    "footer.imprint": "Mentions légales",
    "footer.gdpr": "RGPD",
    "footer.terms": "Conditions d'utilisation",
    "footer.brandName": "4nuel",
    "language.choose": "Choisir la langue",
    "mode.dark": "Sombre",
    "mode.light": "Clair",
    "mode.auto": "Auto",
  },
  // Simplified versions for other languages
  es: {
    "nav.offer": "Oferta",
    "nav.pricing": "Precios",
    "nav.qa": "Preguntas",
    "hero.title": "Soluciones Profesionales de Navegador en la Nube",
    "hero.subtitle": "Proporcionamos tecnología de navegador en la nube de vanguardia para empresas y desarrolladores.",
    "hero.cta": "Contactar",
    "hero.maskTitle": "Infraestructura de Navegador Avanzada",
    "hero.maskSubtitle": "Experimente el futuro de la computación en la nube.",
    "pricing.title": "Precios",
    "pricing.subtitle": "Elija el plan que mejor se adapte a sus necesidades",
    "pricing.mostPopular": "MÁS POPULAR",
    "pricing.hourly.name": "Por Hora", "pricing.hourly.price": "100€", "pricing.hourly.period": "/hora", "pricing.hourly.description": "Perfecto para tareas rápidas",
    "pricing.hourly.feature1": "Pago por uso", "pricing.hourly.feature2": "Configuración instantánea", "pricing.hourly.feature3": "Soporte básico", "pricing.hourly.feature4": "Sin compromisos",
    "pricing.daily.name": "Diario", "pricing.daily.price": "800€", "pricing.daily.period": "/día", "pricing.daily.description": "Ideal para proyectos en curso",
    "pricing.daily.feature1": "Acceso todo el día", "pricing.daily.feature2": "Soporte prioritario", "pricing.daily.feature3": "Características avanzadas", "pricing.daily.feature4": "Disponibilidad 24/7", "pricing.daily.feature5": "Configuraciones personalizadas",
    "pricing.startup.name": "Startups", "pricing.startup.price": "Co-fundador", "pricing.startup.period": "", "pricing.startup.description": "Participación en capital",
    "pricing.startup.feature1": "Asociación basada en capital", "pricing.startup.feature2": "Participación activa", "pricing.startup.feature3": "Compromiso a largo plazo", "pricing.startup.feature4": "Propiedad compartida", "pricing.startup.feature5": "Participación estratégica",
    "qa.title": "Preguntas", "qa.subtitle": "Todo lo que necesita saber",
    "qa.services.question": "¿Qué tipos de servicios ofrezco?", "qa.services.answer": "Me especializo en desarrollo web full-stack, diseño de API, arquitectura en la nube y soluciones de software personalizadas.",
    "qa.approach.question": "¿Cómo abordo nuevos proyectos?", "qa.approach.answer": "Comienzo con una fase de descubrimiento exhaustiva para entender sus requerimientos.",
    "qa.pricingInfo.question": "¿Cuál es mi estructura de precios?", "qa.pricingInfo.answer": "Ofrezco modelos de precios flexibles incluyendo tarifas por hora y asociaciones basadas en capital.",
    "qa.technologies.question": "¿Con qué tecnologías trabajo?", "qa.technologies.answer": "Tengo competencia en tecnologías web modernas incluyendo React, Next.js, Node.js y TypeScript.",
    "qa.quality.question": "¿Cómo aseguro la calidad del código?", "qa.quality.answer": "Sigo las mejores prácticas incluyendo desarrollo dirigido por pruebas y revisiones de código.",
    "footer.imprint": "Aviso Legal", "footer.gdpr": "RGPD", "footer.terms": "Términos de Servicio", "footer.brandName": "4nuel",
    "language.choose": "Elegir Idioma", "mode.dark": "Oscuro", "mode.light": "Claro", "mode.auto": "Auto",
  },
  // Add minimal entries for other languages to prevent errors
  it: { "nav.offer": "Offerta", "nav.pricing": "Prezzi", "nav.qa": "Domande", "hero.title": "Soluzioni Browser Cloud Professionali", "hero.subtitle": "Forniamo tecnologia browser cloud all'avanguardia.", "hero.cta": "Contattaci", "hero.maskTitle": "Infrastruttura Browser Avanzata", "hero.maskSubtitle": "Sperimenta il futuro del cloud computing.", "pricing.title": "Prezzi", "pricing.subtitle": "Scegli il piano migliore", "pricing.mostPopular": "PIÙ POPOLARE", "pricing.hourly.name": "Orario", "pricing.hourly.price": "90€", "pricing.hourly.period": "/ora", "pricing.hourly.description": "Perfetto per compiti rapidi", "pricing.hourly.feature1": "Pagamento per utilizzo", "pricing.hourly.feature2": "Configurazione istantanea", "pricing.hourly.feature3": "Supporto base", "pricing.hourly.feature4": "Nessun impegno", "pricing.daily.name": "Giornaliero", "pricing.daily.price": "720€", "pricing.daily.period": "/giorno", "pricing.daily.description": "Ideale per progetti in corso", "pricing.daily.feature1": "Accesso tutto il giorno", "pricing.daily.feature2": "Supporto prioritario", "pricing.daily.feature3": "Funzionalità avanzate", "pricing.daily.feature4": "Disponibilità 24/7", "pricing.daily.feature5": "Configurazioni personalizzate", "pricing.startup.name": "Startups", "pricing.startup.price": "Co-fondatore", "pricing.startup.period": "", "pricing.startup.description": "Partecipazione azionaria", "pricing.startup.feature1": "Partnership basata su equity", "pricing.startup.feature2": "Partecipazione attiva", "pricing.startup.feature3": "Impegno a lungo termine", "pricing.startup.feature4": "Proprietà condivisa", "pricing.startup.feature5": "Coinvolgimento strategico", "qa.title": "Domande", "qa.subtitle": "Tutto quello che devi sapere", "qa.services.question": "Che tipi di servizi offro?", "qa.services.answer": "Mi specializzo nello sviluppo web full-stack.", "qa.approach.question": "Come affronto nuovi progetti?", "qa.approach.answer": "Inizio con una fase di scoperta approfondita.", "qa.pricingInfo.question": "Qual è la mia struttura dei prezzi?", "qa.pricingInfo.answer": "Offro modelli di prezzo flessibili.", "qa.technologies.question": "Con quali tecnologie lavoro?", "qa.technologies.answer": "Sono competente nelle moderne tecnologie web.", "qa.quality.question": "Come assicuro la qualità del codice?", "qa.quality.answer": "Seguo le migliori pratiche del settore.", "footer.imprint": "Note legali", "footer.gdpr": "GDPR", "footer.terms": "Termini di Servizio", "footer.brandName": "4nuel", "language.choose": "Scegli Lingua", "mode.dark": "Scuro", "mode.light": "Chiaro", "mode.auto": "Auto" },
  pt: { "nav.offer": "Oferta", "nav.pricing": "Preços", "nav.qa": "Perguntas", "hero.title": "Soluções Profissionais de Navegador em Nuvem", "hero.subtitle": "Fornecemos tecnologia de navegador em nuvem de ponta.", "hero.cta": "Entre em Contato", "hero.maskTitle": "Infraestrutura de Navegador Avançada", "hero.maskSubtitle": "Experimente o futuro da computação em nuvem.", "pricing.title": "Preços", "pricing.subtitle": "Escolha o plano que melhor se adapta", "pricing.mostPopular": "MAIS POPULAR", "pricing.hourly.name": "Por Hora", "pricing.hourly.price": "R$500", "pricing.hourly.period": "/hora", "pricing.hourly.description": "Perfeito para tarefas rápidas", "pricing.hourly.feature1": "Pagamento por uso", "pricing.hourly.feature2": "Configuração instantânea", "pricing.hourly.feature3": "Suporte básico", "pricing.hourly.feature4": "Sem compromissos", "pricing.daily.name": "Diário", "pricing.daily.price": "R$4000", "pricing.daily.period": "/dia", "pricing.daily.description": "Ideal para projetos em andamento", "pricing.daily.feature1": "Acesso o dia todo", "pricing.daily.feature2": "Suporte prioritário", "pricing.daily.feature3": "Recursos avançados", "pricing.daily.feature4": "Disponibilidade 24/7", "pricing.daily.feature5": "Configurações personalizadas", "pricing.startup.name": "Startups", "pricing.startup.price": "Co-fundador", "pricing.startup.period": "", "pricing.startup.description": "Participação acionária", "pricing.startup.feature1": "Parceria baseada em equity", "pricing.startup.feature2": "Participação ativa", "pricing.startup.feature3": "Compromisso de longo prazo", "pricing.startup.feature4": "Propriedade compartilhada", "pricing.startup.feature5": "Envolvimento estratégico", "qa.title": "Perguntas", "qa.subtitle": "Tudo o que você precisa saber", "qa.services.question": "Que tipos de serviços ofereço?", "qa.services.answer": "Especializo-me em desenvolvimento web full-stack.", "qa.approach.question": "Como abordo novos projetos?", "qa.approach.answer": "Começo com uma fase de descoberta completa.", "qa.pricingInfo.question": "Qual é minha estrutura de preços?", "qa.pricingInfo.answer": "Ofereço modelos de preços flexíveis.", "qa.technologies.question": "Com quais tecnologias trabalho?", "qa.technologies.answer": "Sou proficiente em tecnologias web modernas.", "qa.quality.question": "Como garanto a qualidade do código?", "qa.quality.answer": "Sigo as melhores práticas da indústria.", "footer.imprint": "Dados legais", "footer.gdpr": "LGPD", "footer.terms": "Termos de Serviço", "footer.brandName": "4nuel", "language.choose": "Escolher Idioma", "mode.dark": "Escuro", "mode.light": "Claro", "mode.auto": "Auto" },
  nl: { "nav.offer": "Aanbod", "nav.pricing": "Prijzen", "nav.qa": "Vragen", "hero.title": "Professionele Cloud Browser Oplossingen", "hero.subtitle": "We bieden geavanceerde cloud browser technologie.", "hero.cta": "Neem Contact Op", "hero.maskTitle": "Geavanceerde Browser Infrastructuur", "hero.maskSubtitle": "Ervaar de toekomst van cloud computing.", "pricing.title": "Prijzen", "pricing.subtitle": "Kies het plan dat het beste past", "pricing.mostPopular": "MEEST POPULAIR", "pricing.hourly.name": "Per Uur", "pricing.hourly.price": "€90", "pricing.hourly.period": "/uur", "pricing.hourly.description": "Perfect voor snelle taken", "pricing.hourly.feature1": "Betaal per gebruik", "pricing.hourly.feature2": "Directe installatie", "pricing.hourly.feature3": "Basis ondersteuning", "pricing.hourly.feature4": "Geen verplichtingen", "pricing.daily.name": "Dagelijks", "pricing.daily.price": "€720", "pricing.daily.period": "/dag", "pricing.daily.description": "Ideaal voor lopende projecten", "pricing.daily.feature1": "Toegang hele dag", "pricing.daily.feature2": "Prioriteit ondersteuning", "pricing.daily.feature3": "Geavanceerde functies", "pricing.daily.feature4": "24/7 beschikbaarheid", "pricing.daily.feature5": "Aangepaste configuraties", "pricing.startup.name": "Startups", "pricing.startup.price": "Medeoprichter", "pricing.startup.period": "", "pricing.startup.description": "Eigenkapitaal participatie", "pricing.startup.feature1": "Op eigenkapital gebaseerd partnerschap", "pricing.startup.feature2": "Actieve deelname", "pricing.startup.feature3": "Lange termijn commitment", "pricing.startup.feature4": "Gedeeld eigendom", "pricing.startup.feature5": "Strategische betrokkenheid", "qa.title": "Vragen", "qa.subtitle": "Alles wat je moet weten", "qa.services.question": "Welke diensten bied ik aan?", "qa.services.answer": "Ik specialiseer me in full-stack webontwikkeling.", "qa.approach.question": "Hoe benader ik nieuwe projecten?", "qa.approach.answer": "Ik begin met een grondige ontdekkingsfase.", "qa.pricingInfo.question": "Wat is mijn prijsstructuur?", "qa.pricingInfo.answer": "Ik bied flexibele prijsmodellen aan.", "qa.technologies.question": "Met welke technologieën werk ik?", "qa.technologies.answer": "Ik ben bedreven in moderne webtechnologieën.", "qa.quality.question": "Hoe zorg ik voor codekwaliteit?", "qa.quality.answer": "Ik volg de beste praktijken van de industrie.", "footer.imprint": "Colofon", "footer.gdpr": "AVG", "footer.terms": "Servicevoorwaarden", "footer.brandName": "4nuel", "language.choose": "Kies Taal", "mode.dark": "Donker", "mode.light": "Licht", "mode.auto": "Auto" },
  ru: { "nav.offer": "Предложение", "nav.pricing": "Цены", "nav.qa": "Вопросы", "hero.title": "Профессиональные облачные браузерные решения", "hero.subtitle": "Мы предоставляем передовые облачные браузерные технологии.", "hero.cta": "Связаться", "hero.maskTitle": "Продвинутая браузерная инфраструктура", "hero.maskSubtitle": "Испытайте будущее облачных вычислений.", "pricing.title": "Цены", "pricing.subtitle": "Выберите план, который лучше всего подходит", "pricing.mostPopular": "САМЫЙ ПОПУЛЯРНЫЙ", "pricing.hourly.name": "Почасовой", "pricing.hourly.price": "₽9000", "pricing.hourly.period": "/час", "pricing.hourly.description": "Идеально для быстрых задач", "pricing.hourly.feature1": "Оплата по использованию", "pricing.hourly.feature2": "Мгновенная настройка", "pricing.hourly.feature3": "Базовая поддержка", "pricing.hourly.feature4": "Без обязательств", "pricing.daily.name": "Ежедневный", "pricing.daily.price": "₽72000", "pricing.daily.period": "/день", "pricing.daily.description": "Идеально для текущих проектов", "pricing.daily.feature1": "Доступ весь день", "pricing.daily.feature2": "Приоритетная поддержка", "pricing.daily.feature3": "Расширенные функции", "pricing.daily.feature4": "Доступность 24/7", "pricing.daily.feature5": "Пользовательские конфигурации", "pricing.startup.name": "Стартапы", "pricing.startup.price": "Соучредитель", "pricing.startup.period": "", "pricing.startup.description": "Участие в капитале", "pricing.startup.feature1": "Партнерство на основе капитала", "pricing.startup.feature2": "Активное участие", "pricing.startup.feature3": "Долгосрочные обязательства", "pricing.startup.feature4": "Совместная собственность", "pricing.startup.feature5": "Стратегическое участие", "qa.title": "Вопросы", "qa.subtitle": "Всё, что нужно знать", "qa.services.question": "Какие услуги я предлагаю?", "qa.services.answer": "Я специализируюсь на full-stack веб-разработке.", "qa.approach.question": "Как я подхожу к новым проектам?", "qa.approach.answer": "Я начинаю с тщательной фазы исследования.", "qa.pricingInfo.question": "Какова моя структура цен?", "qa.pricingInfo.answer": "Я предлагаю гибкие модели ценообразования.", "qa.technologies.question": "С какими технологиями я работаю?", "qa.technologies.answer": "Я владею современными веб-технологиями.", "qa.quality.question": "Как я обеспечиваю качество кода?", "qa.quality.answer": "Я следую лучшим практикам индустрии.", "footer.imprint": "Выходные данные", "footer.gdpr": "GDPR", "footer.terms": "Условия обслуживания", "footer.brandName": "4nuel", "language.choose": "Выбрать язык", "mode.dark": "Тёмный", "mode.light": "Светлый", "mode.auto": "Авто" },
  ja: { "nav.offer": "サービス", "nav.pricing": "料金", "nav.qa": "Q&A", "hero.title": "プロフェッショナルクラウドブラウザソリューション", "hero.subtitle": "最先端のクラウドブラウザ技術を提供します。", "hero.cta": "お問い合わせ", "hero.maskTitle": "高度なブラウザインフラストラクチャ", "hero.maskSubtitle": "クラウドコンピューティングの未来を体験してください。", "pricing.title": "料金", "pricing.subtitle": "最適なプランをお選びください", "pricing.mostPopular": "最も人気", "pricing.hourly.name": "時間制", "pricing.hourly.price": "¥15,000", "pricing.hourly.period": "/時間", "pricing.hourly.description": "短期タスクに最適", "pricing.hourly.feature1": "従量課金", "pricing.hourly.feature2": "即座セットアップ", "pricing.hourly.feature3": "基本サポート", "pricing.hourly.feature4": "契約不要", "pricing.daily.name": "日額制", "pricing.daily.price": "¥120,000", "pricing.daily.period": "/日", "pricing.daily.description": "継続的なプロジェクトに理想的", "pricing.daily.feature1": "終日アクセス", "pricing.daily.feature2": "優先サポート", "pricing.daily.feature3": "高度な機能", "pricing.daily.feature4": "24/7可用性", "pricing.daily.feature5": "カスタム設定", "pricing.startup.name": "スタートアップ", "pricing.startup.price": "共同創設者", "pricing.startup.period": "", "pricing.startup.description": "エクイティ参加", "pricing.startup.feature1": "エクイティベースパートナーシップ", "pricing.startup.feature2": "積極的参加", "pricing.startup.feature3": "長期コミット", "pricing.startup.feature4": "共有所有権", "pricing.startup.feature5": "戦略的関与", "qa.title": "よくある質問", "qa.subtitle": "知っておくべき全ての情報", "qa.services.question": "どのようなサービスを提供していますか？", "qa.services.answer": "フルスタックWeb開発を専門としています。", "qa.approach.question": "新しいプロジェクトにどうアプローチしますか？", "qa.approach.answer": "要件を理解するための徹底的な発見段階から始めます。", "qa.pricingInfo.question": "料金体系はどのようになっていますか？", "qa.pricingInfo.answer": "柔軟な料金モデルを提供しています。", "qa.technologies.question": "どの技術を使用していますか？", "qa.technologies.answer": "最新のWeb技術に精通しています。", "qa.quality.question": "コード品質をどのように確保していますか？", "qa.quality.answer": "業界のベストプラクティスに従っています。", "footer.imprint": "法的表示", "footer.gdpr": "GDPR", "footer.terms": "利用規約", "footer.brandName": "4nuel", "language.choose": "言語を選択", "mode.dark": "ダーク", "mode.light": "ライト", "mode.auto": "自動" },
  ko: { "nav.offer": "제안", "nav.pricing": "가격", "nav.qa": "질문", "hero.title": "전문 클라우드 브라우저 솔루션", "hero.subtitle": "최첨단 클라우드 브라우저 기술을 제공합니다.", "hero.cta": "연락하기", "hero.maskTitle": "고급 브라우저 인프라", "hero.maskSubtitle": "클라우드 컴퓨팅의 미래를 경험하세요.", "pricing.title": "가격", "pricing.subtitle": "비즈니스에 가장 적합한 플랜을 선택하세요", "pricing.mostPopular": "가장 인기", "pricing.hourly.name": "시간당", "pricing.hourly.price": "₩135,000", "pricing.hourly.period": "/시간", "pricing.hourly.description": "빠른 작업에 완벽", "pricing.hourly.feature1": "사용한 만큼 지불", "pricing.hourly.feature2": "즉시 설정", "pricing.hourly.feature3": "기본 지원", "pricing.hourly.feature4": "약정 없음", "pricing.daily.name": "일일", "pricing.daily.price": "₩1,080,000", "pricing.daily.period": "/일", "pricing.daily.description": "진행 중인 프로젝트에 이상적", "pricing.daily.feature1": "하루 종일 액세스", "pricing.daily.feature2": "우선 지원", "pricing.daily.feature3": "고급 기능", "pricing.daily.feature4": "24/7 가용성", "pricing.daily.feature5": "맞춤 구성", "pricing.startup.name": "스타트업", "pricing.startup.price": "공동 창립자", "pricing.startup.period": "", "pricing.startup.description": "지분 참여", "pricing.startup.feature1": "지분 기반 파트너십", "pricing.startup.feature2": "적극적 참여", "pricing.startup.feature3": "장기 약속", "pricing.startup.feature4": "공유 소유권", "pricing.startup.feature5": "전략적 참여", "qa.title": "질문", "qa.subtitle": "함께 일하는 것에 대해 알아야 할 모든 것", "qa.services.question": "어떤 종류의 서비스를 제공하나요?", "qa.services.answer": "풀스택 웹 개발을 전문으로 합니다.", "qa.approach.question": "새로운 프로젝트에 어떻게 접근하나요?", "qa.approach.answer": "요구 사항을 이해하기 위한 철저한 발견 단계부터 시작합니다.", "qa.pricingInfo.question": "가격 구조는 어떻게 되나요?", "qa.pricingInfo.answer": "유연한 가격 모델을 제공합니다.", "qa.technologies.question": "어떤 기술을 사용하나요?", "qa.technologies.answer": "최신 웹 기술에 능숙합니다.", "qa.quality.question": "코드 품질을 어떻게 보장하나요?", "qa.quality.answer": "업계 모범 사례를 따릅니다.", "footer.imprint": "법적 고지", "footer.gdpr": "GDPR", "footer.terms": "서비스 약관", "footer.brandName": "4nuel", "language.choose": "언어 선택", "mode.dark": "다크", "mode.light": "라이트", "mode.auto": "자동" },
  zh: { "nav.offer": "服务", "nav.pricing": "价格", "nav.qa": "问答", "hero.title": "专业云浏览器解决方案", "hero.subtitle": "我们为企业和开发者提供前沿的云浏览器技术。", "hero.cta": "联系我们", "hero.maskTitle": "先进的浏览器基础设施", "hero.maskSubtitle": "体验云计算的未来。", "pricing.title": "价格", "pricing.subtitle": "选择最适合您业务需求的计划", "pricing.mostPopular": "最受欢迎", "pricing.hourly.name": "按小时", "pricing.hourly.price": "¥720", "pricing.hourly.period": "/小时", "pricing.hourly.description": "适合快速任务", "pricing.hourly.feature1": "按用量付费", "pricing.hourly.feature2": "即时设置", "pricing.hourly.feature3": "基础支持", "pricing.hourly.feature4": "无承诺", "pricing.daily.name": "按天", "pricing.daily.price": "¥5,760", "pricing.daily.period": "/天", "pricing.daily.description": "适合持续项目", "pricing.daily.feature1": "全天访问", "pricing.daily.feature2": "优先支持", "pricing.daily.feature3": "高级功能", "pricing.daily.feature4": "24/7可用性", "pricing.daily.feature5": "自定义配置", "pricing.startup.name": "初创公司", "pricing.startup.price": "联合创始人", "pricing.startup.period": "", "pricing.startup.description": "股权参与", "pricing.startup.feature1": "基于股权的合作伙伴关系", "pricing.startup.feature2": "积极参与", "pricing.startup.feature3": "长期承诺", "pricing.startup.feature4": "共享所有权", "pricing.startup.feature5": "战略参与", "qa.title": "常见问题", "qa.subtitle": "关于合作需要了解的一切", "qa.services.question": "我提供什么类型的服务？", "qa.services.answer": "我专门从事全栈Web开发。", "qa.approach.question": "我如何处理新项目？", "qa.approach.answer": "我从彻底的发现阶段开始，以了解您的需求。", "qa.pricingInfo.question": "我的定价结构是什么？", "qa.pricingInfo.answer": "我提供灵活的定价模型。", "qa.technologies.question": "我使用什么技术？", "qa.technologies.answer": "我精通现代Web技术。", "qa.quality.question": "我如何确保代码质量？", "qa.quality.answer": "我遵循行业最佳实践。", "footer.imprint": "版权信息", "footer.gdpr": "GDPR", "footer.terms": "服务条款", "footer.brandName": "4nuel", "language.choose": "选择语言", "mode.dark": "深色", "mode.light": "浅色", "mode.auto": "自动" },
  ar: { "nav.offer": "العرض", "nav.pricing": "الأسعار", "nav.qa": "الأسئلة", "hero.title": "حلول متصفح سحابي احترافية", "hero.subtitle": "نوفر تقنية متصفح سحابي متطورة للشركات والمطورين.", "hero.cta": "تواصل معنا", "hero.maskTitle": "بنية تحتية متقدمة للمتصفح", "hero.maskSubtitle": "اختبر مستقبل الحوسبة السحابية.", "pricing.title": "الأسعار", "pricing.subtitle": "اختر الخطة الأنسب لاحتياجات عملك", "pricing.mostPopular": "الأكثر شيوعاً", "pricing.hourly.name": "بالساعة", "pricing.hourly.price": "$100", "pricing.hourly.period": "/ساعة", "pricing.hourly.description": "مثالي للمهام السريعة", "pricing.hourly.feature1": "ادفع حسب الاستخدام", "pricing.hourly.feature2": "إعداد فوري", "pricing.hourly.feature3": "دعم أساسي", "pricing.hourly.feature4": "بدون التزامات", "pricing.daily.name": "يومي", "pricing.daily.price": "$800", "pricing.daily.period": "/يوم", "pricing.daily.description": "مثالي للمشاريع الجارية", "pricing.daily.feature1": "وصول طوال اليوم", "pricing.daily.feature2": "دعم أولوية", "pricing.daily.feature3": "ميزات متقدمة", "pricing.daily.feature4": "توفر 24/7", "pricing.daily.feature5": "تكوينات مخصصة", "pricing.startup.name": "الشركات الناشئة", "pricing.startup.price": "شريك مؤسس", "pricing.startup.period": "", "pricing.startup.description": "مشاركة في رأس المال", "pricing.startup.feature1": "شراكة قائمة على رأس المال", "pricing.startup.feature2": "مشاركة نشطة", "pricing.startup.feature3": "التزام طويل الأمد", "pricing.startup.feature4": "ملكية مشتركة", "pricing.startup.feature5": "مشاركة استراتيجية", "qa.title": "الأسئلة", "qa.subtitle": "كل ما تحتاج لمعرفته حول العمل معاً", "qa.services.question": "ما أنواع الخدمات التي أقدمها؟", "qa.services.answer": "أتخصص في تطوير الويب الكامل.", "qa.approach.question": "كيف أتعامل مع المشاريع الجديدة؟", "qa.approach.answer": "أبدأ بمرحلة اكتشاف شاملة لفهم متطلباتك.", "qa.pricingInfo.question": "ما هي هيكل أسعاري؟", "qa.pricingInfo.answer": "أقدم نماذج تسعير مرنة.", "qa.technologies.question": "ما التقنيات التي أعمل بها؟", "qa.technologies.answer": "أجيد التقنيات الحديثة للويب.", "qa.quality.question": "كيف أضمن جودة الكود؟", "qa.quality.answer": "أتبع أفضل الممارسات في الصناعة.", "footer.imprint": "بيانات قانونية", "footer.gdpr": "GDPR", "footer.terms": "شروط الخدمة", "footer.brandName": "4nuel", "language.choose": "اختيار اللغة", "mode.dark": "داكن", "mode.light": "فاتح", "mode.auto": "تلقائي" }
}

interface I18nProviderProps {
  children: ReactNode
}

export function I18nProvider({ children }: I18nProviderProps) {
  const [locale, setLocaleState] = useState<Locale>("en")

  useEffect(() => {
    const savedLocale = localStorage.getItem("locale") as Locale
    if (savedLocale && translations[savedLocale]) {
      setLocaleState(savedLocale)
    }
  }, [])

  const setLocale = (newLocale: Locale) => {
    setLocaleState(newLocale)
    localStorage.setItem("locale", newLocale)
  }

  const t = (key: string): string => {
    return translations[locale]?.[key] || translations.en[key] || key
  }

  const contextValue: I18nContextType = {
    locale,
    setLocale,
    t,
  }

  return (
    <I18nContext.Provider value={contextValue}>
      {children}
    </I18nContext.Provider>
  )
}

export function useI18n(): I18nContextType {
  const context = useContext(I18nContext)
  if (!context) {
    throw new Error("useI18n must be used within an I18nProvider")
  }
  return context
}