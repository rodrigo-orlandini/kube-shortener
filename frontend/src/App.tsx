import { useState, useEffect } from 'react'
import './App.css'
import { BarChart, Bar, XAxis, YAxis, Tooltip, ResponsiveContainer, CartesianGrid } from 'recharts';

interface AnalyticsData {
  shortUrl: string;
  visitCount: number;
}

interface AnalyticsResponse {
  status: string;
  data: AnalyticsData[];
}

function App() {
  const [url, setUrl] = useState('');
  const [shortUrl, setShortUrl] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const [analytics, setAnalytics] = useState<AnalyticsData[]>([]);

  const handleShorten = async () => {
    setLoading(true);
    setError('');
    setShortUrl('');

    try {
      const response = await fetch('http://localhost:8080/shorten', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ originalUrl: url }),
      });

      if (!response.ok) throw new Error('Failed to shorten URL');

      const data = await response.json();
      const shortCode = data.shortenedUrl || '';
      setShortUrl(`http://localhost:8080${shortCode}`);
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : 'Unknown error');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    const fetchAnalytics = async () => {
      try {
        const response = await fetch('http://localhost:8081/visits/topRanked?limit=5');
        if (!response.ok) {
          throw new Error('Failed to fetch analytics');
        }
        const data: AnalyticsResponse = await response.json();
        setAnalytics(data.data || []);
      } catch (err) {
        console.error('Error fetching analytics:', err);
        setAnalytics([]);
      }
    };

    fetchAnalytics();
  }, []);

  const formatUrlForDisplay = (shortUrl: string) => {
    const urlParts = shortUrl.split('/');
    return urlParts[urlParts.length - 1] || shortUrl;
  };

  return (
    <div className="shortener-container">
      <h1 className="shortener-title">URL Shortener</h1>
      <p className="shortener-subtitle">Transform long URLs into short, shareable links instantly</p>

      <div className="shortener-form">
        <input
          className="shortener-input"
          type="url"
          placeholder="Paste your long URL here..."
          value={url}
          onChange={e => setUrl(e.target.value)}
          disabled={loading}
        />
        <button
          className={`shortener-button${loading ? ' loading' : ''}`}
          onClick={handleShorten}
          disabled={!url || loading}
        >
          {loading ? 'Shortening...' : 'Shorten'}
        </button>
      </div>

      {shortUrl && (
        <div className="shortener-result">
          <span>Your short URL:</span>
          <a href={shortUrl} target="_blank" rel="noopener noreferrer">{shortUrl}</a>
        </div>
      )}
      {error && <div className="shortener-error">{error}</div>}
			
      <div className="shortener-analytics">
        <h2 className="shortener-analytics-title">Top 5 Most Visited Shortened URLs</h2>
        {analytics.length > 0 ? (
          <ResponsiveContainer width="100%" height={320}>
            <BarChart data={analytics} margin={{ top: 16, right: 24, left: 0, bottom: 8 }}>
              <CartesianGrid strokeDasharray="3 3" stroke="#222" />
              <XAxis 
                dataKey="shortUrl" 
                stroke="#aaa" 
                tick={{ fill: '#aaa', fontSize: 12 }} 
                interval={0} 
                angle={-20} 
                textAnchor="end" 
                height={60}
                tickFormatter={formatUrlForDisplay}
              />
              <YAxis stroke="#aaa" tick={{ fill: '#aaa', fontSize: 12 }} allowDecimals={false} />
              <Tooltip 
                contentStyle={{ background: '#181a20', border: 'none', color: '#fff' }} 
                labelStyle={{ color: '#aaa' }}
                formatter={(value: number) => [value, 'Visits']}
                labelFormatter={(label: string) => `URL: ${formatUrlForDisplay(label)}`}
              />
              <Bar dataKey="visitCount" fill="#646cff" radius={[8, 8, 0, 0]} />
            </BarChart>
          </ResponsiveContainer>
        ) : (
          <div className="shortener-analytics-empty">
            <p>No analytics data available</p>
          </div>
        )}
      </div>
    </div>
  );
}

export default App
