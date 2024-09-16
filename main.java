public static String[] extractNames(String text) {
        // Регулярное выражение для поиска имен
        String regexWoman = "\\b[А-Я][а-я]+ова\\s([А-Я][а-я]+)\\s[А-Я][а-я]+овна\\b";
        String regexMan = "\\b[А-Я][а-я]+ов\\s[А-Я][а-я]+\\s[А-Я][а-я]+ович";
        String regexI = "([А-Я][а-я]+ов|[А-Я][а-я]+ова)\\s[А-Я]\\.[А-Я]\\.";
        List<String> names = new ArrayList<>();
        
        
        Pattern pattern = Pattern.compile(regexMan);
        Matcher matcher = pattern.matcher(text);
        
        while (matcher.find()) {
            String fullName = matcher.group(0);
            names.add(fullName);
        }
        
        pattern = Pattern.compile(regexI);
        matcher = pattern.matcher(text);
        while (matcher.find()) {
            String fullName = matcher.group(0);
            names.add(fullName);
        } 
        pattern = Pattern.compile(regexWoman);
        matcher = pattern.matcher(text);
        while (matcher.find()) {
            String fullName = matcher.group(0);
            names.add(fullName);
        }
        return names.toArray(new String[0]);
    }
   
