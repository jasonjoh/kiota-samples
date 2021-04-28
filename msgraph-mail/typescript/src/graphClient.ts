import {HttpCore, HttpMethod, RequestInfo, ResponseHandler, SerializationWriterFactory} from '@microsoft/kiota-abstractions';
import {UsersRequestBuilder} from './users/usersRequestBuilder';

/** The main entry point of the SDK, exposes the configuration and the fluent API.  */
export class GraphClient {
    /** Current path for the request  */
    public currentPath?: string | undefined;
    /** Core service to use to execute the requests  */
    public httpCore?: HttpCore | undefined;
    /** Path segment to use to build the URL for the current request builder  */
    private readonly pathSegment: string = "https://graph.microsoft.com/v1.0";
    /** Factory to use to get a serializer for payload serialization  */
    public serializerFactory?: SerializationWriterFactory | undefined;
    public get users(): UsersRequestBuilder {
        const builder = new UsersRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    }
}
