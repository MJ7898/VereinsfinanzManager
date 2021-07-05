# Here you can see the Architecture of our Go-Backend

The following link is to our overall architecture diagram to get some insights of our application and our minds.
[application_diagramm](../Diagrams/application_architecture.drawio)

## Ausgangsituation 
Das System ist primär in drei Bereiche unterteilt:
1. Club-Ebene
2. Abteilungsebene
3. Team- / Mannschaftsebene

### Absichten
#### Club
Es wurde sich dazu entschieden, den Club mit grundlegenden Informationen erstellen zu können. 
Hinzu kommt, dass der Club über ein Budget und Kontingent verfügt, welches quer über 
den ganzen Verein verteilt werden kann. Die Ausgangsposition bezüglich des Geldes ist jedoch immer der komplette Verein. 
Bei der Anlage eines Vereins, können nachträglich Abteilungen erstellt werden, die im Verein auf oberste Ebene mit der Identifikationsnummer registriert werden. 
Dahingehend kann der CLub lediglich auf seine Abteilungen vergrößert oder verkleinert (aktualisiert) werden. 
Beim Löschen eines Vereins werden sämtliche Abteilungen und Teams mitgelöscht. Es werden keine Informationen (Datenleichen) liegengelassen.
#### Abteilung
Die Abteilung ist von der Einstufung als auch von der Ebene das Herzstück. Diese Ebene beinhaltet die meisten Abhängigkeiten auf die Ebene Club und die darunterliegende Ebene der Teams. 
Die Abteilungen werden im Club registriert. Anders werden die Teams mit der jeweiligen Identifikationsnummer in der Abteilung hinterlegt, damit der READ-Aufwand auf mehrere Collections der NoSQL Datenbank vermieden werden. 
Durch die klare Zuordnung wird deutlich, welche Teams zu einer Abteilung gehören. 
Das Budget einer Abteilung wird über die Teams hinweg desammelt (computed) und auf der Ebene der Abteilung ERmittelt. Hierfür gibt es ein Pattern der NoSQL-Datenbank mongoDB. (Dies ist ein Entscheidungsgrund, wieso in diesem Fall mongoDB eingesetzt wird.)

##### Entscheidung der DB
Die Gründe einer NoSQL-Datenbank werden während der Vorstellung erläutert und können in der Dokumentation unter folgendem Dokument eingesehen werden: 
[Präsentation und Entscheidung auf eine NoSQL-Datenbank mit mongDB](../../Presentation/presentation.pptx)

#### Team
Ein Team enthält Informationen über den aktuellen Zustand des Teams von Budget, Ausgaben etc.

Änderungen im Team betrifft lediglich das Team. Werden Änderungen am Budget vorgenommen, so wird das über die Ebenen hinweg kommuniziert, sodass eine stetige Informationsgabe der Finanzen stattfindet. 
Wird innerhalb einer Abteilung ein neues Team erstellt, so wird die jeweilige Abteilung auf dieses Team ergänzt. Im Gegenzug wird die Abteilung auf das wegfallende Team gemindert.  
In dieser Ebene wird ebenso der darüberliegenden Ebene eine Liste von anhängenden Teams mitgegeben, sodass die allgemeinen Leseoperationen auf mehrere Collections wegfällt. So können Kalkulationen und Analysen deutlich effektiver umgesetzt werden. 