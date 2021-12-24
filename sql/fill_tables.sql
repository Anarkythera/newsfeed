-- Inserting data into sources table

INSERT INTO
    sources (url, provider, category)
VALUES
    ('http://feeds.bbci.co.uk/news/uk/rss.xml', 'BBC News', 'UK'),
    ('http://feeds.skynews.com/feeds/rss/uk.xml','SkyNews', 'UK');
