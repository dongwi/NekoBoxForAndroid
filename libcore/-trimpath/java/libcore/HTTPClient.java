// Code generated by gobind. DO NOT EDIT.

// Java class libcore.HTTPClient is a proxy for talking to a Go program.
//
//   autogenerated by gobind -lang=java libcore
package libcore;

import go.Seq;

public interface HTTPClient {
	public void close();
	public void keepAlive();
	public void modernTLS();
	public HTTPRequest newRequest();
	public void pinnedSHA256(String sumHex);
	public void pinnedTLS12();
	public void restrictedTLS();
	public void trySocks5(int port);
	
}

